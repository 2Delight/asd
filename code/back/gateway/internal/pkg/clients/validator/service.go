package validator

import (
	"context"
	"gateway-api/internal/pkg/model"
	"gateway-api/pkg/validator"
	"google.golang.org/grpc"
)

type Client struct {
	cli validator.ValidatorServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	cli := validator.NewValidatorServiceClient(conn)

	return &Client{
		cli: cli,
	}
}

func (c *Client) ValidateSpecification(ctx context.Context, id int64, specContent string) (*model.ValidationData, error) {
	spec, err := c.cli.ValidateSpecification(ctx, &validator.ValidateSpecificationRequest{
		Specification: &validator.Specification{
			Id:      id,
			Content: specContent,
		},
	})
	if err != nil {
		return nil, err
	}

	validationData := &model.ValidationData{
		IsValid: spec.IsValid,
		Errors:  make([]*model.Error, 0, len(spec.Errors)),
		Hints:   make([]*model.Hint, 0, len(spec.Hints)),
	}
	for _, er := range spec.Errors {
		validationData.Errors = append(validationData.Errors, &model.Error{
			Code:    er.Code,
			Message: er.Message,
		})
	}
	for _, hn := range spec.Hints {
		validationData.Hints = append(validationData.Hints, &model.Hint{
			Message: hn.Message,
		})
	}

	return validationData, nil
}
