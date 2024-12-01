package integrator

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gateway-api/internal/pkg/model"
	"gateway-api/pkg/integrator"
	"google.golang.org/grpc"
)

type Client struct {
	cli integrator.WorkerServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	cli := integrator.NewWorkerServiceClient(conn)

	return &Client{
		cli: cli,
	}
}

func (c *Client) GetSpecificationFromGit(ctx context.Context, id int64) (*model.Specification, error) {
	spec, err := c.cli.GetSpecificationFromGit(ctx, &integrator.GetSpecificationRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &model.Specification{
		Id:        spec.GetId(),
		Name:      spec.GetName(),
		Content:   spec.GetContent(),
		GitPath:   spec.GetGitPath(),
		Status:    spec.GetStatus(),
		CreatedAt: spec.GetCreatedAt().AsTime(),
		UpdatedAt: spec.GetUpdatedAt().AsTime(),
	}, nil
}

func (c *Client) SaveSpecification(ctx context.Context, spec *model.Specification) (*model.PushResult, error) {
	pushRes, err := c.cli.SaveSpecification(ctx, &integrator.SaveSpecificationRequest{
		Id: spec.Id,
		Specification: &integrator.Specification{
			Id:        spec.Id,
			Name:      spec.Name,
			Content:   spec.Content,
			GitPath:   spec.GitPath,
			Status:    spec.Status,
			CreatedAt: timestamppb.New(spec.CreatedAt),
			UpdatedAt: timestamppb.New(spec.UpdatedAt),
		},
	})
	if err != nil {
		return nil, err
	}

	return &model.PushResult{
		CommitHash: pushRes.CommitHash,
		IsSuccess:  pushRes.IsSuccess,
	}, nil
}

func (c *Client) RunMLDev(ctx context.Context, spec *model.Specification) (*model.MLDevResult, error) {
	mlRes, err := c.cli.RunMLDev(ctx, &integrator.RunMLDevRequest{
		Id:   spec.Id,
		Path: &integrator.MLDevPath{Path: spec.GitPath},
	})
	if err != nil {
		return nil, err
	}

	return &model.MLDevResult{
		Artifacts: mlRes.GetArtifacts(),
		IsSuccess: mlRes.IsSuccess,
	}, nil
}
