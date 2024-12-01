package gateway

import (
	"context"

	"gateway-api/internal/pkg/model"
	pbgateway "gateway-api/pkg/gateway"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SpecificationService interface {
	GetSpecification(ctx context.Context, id int64) (*model.Specification, error)
	UpdateSpecification(ctx context.Context, id int64, updatedContent string) (bool, error)
	GetStatus(ctx context.Context, id int64) (*model.Status, error)
	UpdateStatus(ctx context.Context, id int64, updatedStatus string) (bool, error)
	ValidateSpecification(ctx context.Context, id int64, specContent string) (*model.ValidationData, error)
}

type GatewayHandler struct {
	pbgateway.UnimplementedGatewayServiceServer

	specService SpecificationService
}

func NewSpecHandler(specService SpecificationService) *GatewayHandler {
	return &GatewayHandler{
		specService: specService,
	}
}

func (g *GatewayHandler) GetSpecification(ctx context.Context, req *pbgateway.GetSpecificationRequest) (*pbgateway.Specification, error) {
	if req.GetId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id should be greater than 0")
	}

	spec, err := g.specService.GetSpecification(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pbgateway.Specification{
		Id:        spec.Id,
		Name:      spec.Name,
		Content:   spec.Content,
		GitPath:   spec.GitPath,
		Status:    spec.Status,
		CreatedAt: timestamppb.New(spec.CreatedAt),
		UpdatedAt: timestamppb.New(spec.UpdatedAt),
	}, nil
}

func (g *GatewayHandler) UpdateSpecification(ctx context.Context, req *pbgateway.UpdateSpecificationRequest) (*pbgateway.UpdateSpecificationResponse, error) {
	if req.GetId() <= 0 || req.GetSpecificationContent() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id/content")
	}

	isSucc, err := g.specService.UpdateSpecification(ctx, req.GetId(), req.GetSpecificationContent())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pbgateway.UpdateSpecificationResponse{
		IsSuccess: isSucc,
	}, nil
}

func (g *GatewayHandler) GetStatus(ctx context.Context, req *pbgateway.GetStatusRequest) (*pbgateway.Status, error) {
	if req.GetId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id should be greater than 0")
	}

	stat, err := g.specService.GetStatus(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pbgateway.Status{
		Id:   stat.ID,
		Name: stat.Name,
	}, nil
}

func (g *GatewayHandler) UpdateStatus(ctx context.Context, req *pbgateway.UpdateStatusRequest) (*pbgateway.StatusUpdateResponse, error) {
	if req.GetId() <= 0 || (req.GetStatusUpdate() == nil || req.GetStatusUpdate().GetNewStatus() == "") {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id/status")
	}

	isSucc, err := g.specService.UpdateStatus(ctx, req.GetId(), req.GetStatusUpdate().GetNewStatus())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pbgateway.StatusUpdateResponse{
		IsSuccess: isSucc,
	}, nil
}

func (g *GatewayHandler) ValidateSpecification(ctx context.Context, req *pbgateway.ValidateSpecificationRequest) (*pbgateway.ValidationResult, error) {
	if req.GetId() <= 0 || req.GetSpecificationContent() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id/content")
	}

	result, err := g.specService.ValidateSpecification(ctx, req.GetId(), req.GetSpecificationContent())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	resp := &pbgateway.ValidationResult{
		IsValid: result.IsValid,
		Errors:  make([]*pbgateway.Error, 0, len(result.Errors)),
		Hints:   make([]*pbgateway.Hint, 0, len(result.Hints)),
	}
	for _, er := range result.Errors {
		resp.Errors = append(resp.Errors, &pbgateway.Error{
			Code:    er.Code,
			Message: er.Message,
		})
	}
	for _, hn := range result.Hints {
		resp.Hints = append(resp.Hints, &pbgateway.Hint{
			Message: hn.Message,
		})
	}

	return resp, nil
}

func (g *GatewayHandler) GetHello(context.Context, *pbgateway.GetHelloRequest) (*pbgateway.GetHelloResponse, error) {
	return &pbgateway.GetHelloResponse{Pong: "kek"}, nil
}
