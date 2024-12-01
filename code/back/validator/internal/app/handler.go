package validator

import (
	"context"

	pbvalidator "validator-api/pkg/validator"
)

type ValidatorHandler struct {
	pbvalidator.UnimplementedValidatorServiceServer
}

func NewValidatorHandler() *ValidatorHandler {
	return &ValidatorHandler{}
}

func (h *ValidatorHandler) ValidateSpecification(context.Context, *pbvalidator.ValidateSpecificationRequest) (*pbvalidator.ValidationResult, error) {
	return &pbvalidator.ValidationResult{
		IsValid: true,
		Errors: []*pbvalidator.Error{
			{
				Code:    "11",
				Message: "Некорректное значение переменной",
			},
			{
				Code:    "26",
				Message: "Не существует такой команды",
			},
		},
		Hints: []*pbvalidator.Hint{
			{
				Message: "Вы можете определять alias для переменных",
			},
			{
				Message: "Данный способ определения deprecated",
			},
		},
	}, nil
}

func (h *ValidatorHandler) GetHello(context.Context, *pbvalidator.GetHelloRequest) (*pbvalidator.GetHelloResponse, error) {
	return &pbvalidator.GetHelloResponse{
		Pong: "pep",
	}, nil
}
