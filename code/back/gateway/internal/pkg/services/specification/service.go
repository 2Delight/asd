package specification

import (
	"context"

	"gateway-api/internal/pkg/database/schema"
	"gateway-api/internal/pkg/model"
)

type repo interface {
	GetSpecificationByID(ctx context.Context, id int64) (*schema.Specification, error)
	UpdateSpecificationStatus(ctx context.Context, id int64, newStatus string) error
}

type IntegratorClient interface {
	GetSpecificationFromGit(ctx context.Context, id int64) (*model.Specification, error)
	SaveSpecification(ctx context.Context, spec *model.Specification) (*model.PushResult, error)
	RunMLDev(ctx context.Context, spec *model.Specification) (*model.MLDevResult, error)
}

type ValidatorClient interface {
	ValidateSpecification(ctx context.Context, id int64, specContent string) (*model.ValidationData, error)
}

type SpecService struct {
	integrator IntegratorClient
	validator  ValidatorClient
	repo       repo
}

func NewService(integrator IntegratorClient, validator ValidatorClient, repo repo) *SpecService {
	return &SpecService{
		integrator: integrator,
		validator:  validator,
		repo:       repo,
	}
}

func (s *SpecService) GetSpecification(ctx context.Context, id int64) (*model.Specification, error) {
	gitSpec, err := s.integrator.GetSpecificationFromGit(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.Specification{
		Id:      gitSpec.Id,
		Content: gitSpec.Content,
	}, nil
}

func (s *SpecService) UpdateSpecification(ctx context.Context, id int64, updatedContent string) (bool, error) {
	schemaSpec, err := s.repo.GetSpecificationByID(ctx, id)
	if err != nil {
		return false, err
	}

	if schemaSpec == nil {
		return false, nil
	}

	var (
		spec = &model.Specification{
			Id:        schemaSpec.Id,
			Name:      schemaSpec.Name,
			Content:   updatedContent,
			GitPath:   schemaSpec.GitPath,
			Status:    schemaSpec.Status,
			CreatedAt: schemaSpec.CreatedAt,
			UpdatedAt: schemaSpec.UpdatedAt,
		}
	)
	res, err := s.integrator.SaveSpecification(ctx, spec)
	if err != nil {
		return false, err
	}

	if !res.IsSuccess {
		return false, nil
	}

	mlDevRes, err := s.integrator.RunMLDev(ctx, spec)

	return mlDevRes.IsSuccess, nil
}

func (s *SpecService) GetStatus(ctx context.Context, id int64) (*model.Status, error) {
	schemaSpec, err := s.repo.GetSpecificationByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if schemaSpec == nil {
		return nil, nil
	}

	return &model.Status{
		ID:   schemaSpec.Id,
		Name: schemaSpec.Name,
	}, nil
}

func (s *SpecService) UpdateStatus(ctx context.Context, id int64, updatedStatus string) (bool, error) {
	if err := s.repo.UpdateSpecificationStatus(ctx, id, updatedStatus); err != nil {
		return false, err
	}

	return true, nil
}

func (s *SpecService) ValidateSpecification(ctx context.Context, id int64, specContent string) (*model.ValidationData, error) {
	return s.validator.ValidateSpecification(ctx, id, specContent)
}
