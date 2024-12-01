package integrator

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	pbintegrator "integrator-api/pkg/integrator"
	"math/rand"
	"time"
)

type IntegratorHandler struct {
	pbintegrator.UnimplementedWorkerServiceServer
}

func NewIntegratorHandler() *IntegratorHandler {
	return &IntegratorHandler{}
}
func (h *IntegratorHandler) SaveSpecification(context.Context, *pbintegrator.SaveSpecificationRequest) (*pbintegrator.CommitPushResult, error) {
	return &pbintegrator.CommitPushResult{
		CommitHash: GenerateRandomHash(),
		IsSuccess:  true,
	}, nil
}

func (h *IntegratorHandler) RunMLDev(ctx context.Context, req *pbintegrator.RunMLDevRequest) (*pbintegrator.MLDevResult, error) {
	return &pbintegrator.MLDevResult{
		IsSuccess: true,
		Artifacts: []string{
			fmt.Sprintf("launch_uuid: %s", uuid.New().String()),
			fmt.Sprintf("launched_at: %s", time.Now().String()),
			fmt.Sprintf("spec_id: %d", req.GetId()),
			fmt.Sprintf("spec_git_path: %s", req.GetPath()),
		},
	}, nil
}
func (h *IntegratorHandler) GetSpecificationFromGit(ctx context.Context, req *pbintegrator.GetSpecificationRequest) (*pbintegrator.Specification, error) {
	return &pbintegrator.Specification{
		Id:        req.GetId(),
		Content:   "# Licensed under the Apache License: http://www.apache.org/licenses/LICENSE-2.0\n# For details: https://gitlab.com/mlrep/mldev/-/blob/master/NOTICE.md\n\nprepare: &prepare_stage !BasicStage\n  name: prepare\n  params:\n    size: 1\n  inputs:\n    - !path { path: \"./src\" }\n  outputs:\n    - !path { path: \"./data\" }\n  script:\n    - \"python3 src/prepare.py\"\n\n\ntrain: &train_stage !BasicStage\n  name: train\n  params:\n    num_iters: 10\n  inputs:\n    - !path\n      path: \"./data\"\n      files:\n        - \"X_train.pickle\"\n        - \"X_dev.pickle\"\n        - \"X_test.pickle\"\n        - \"y_train.pickle\"\n        - \"y_dev.pickle\"\n        - \"y_test.pickle\"\n  outputs: &model_data\n    - !path\n      path: \"models/default\"\n      files:\n        - \"model.pickle\"\n  script:\n    - \"python3 src/train.py --n ${self.params.num_iters}\"\n\npresent_model: &present_model !BasicStage\n  name: present_model\n  inputs: *model_data\n  outputs:\n    - !path\n      path: \"results/default\"\n  env:\n    MLDEV_MODEL_PATH: ${path(self.inputs[0].path)}\n    RESULTS_PATH: ${self.outputs[0].path}\n  script:\n    - |\n      python3 src/predict.py\n      printf \"=============================\\n\"\n      printf \"Test report:\\n\\n\"\n      cat ${path(self.outputs[0].path)}/test_report.json\n      printf \"\\n\\n=============================\\n\"\n\npipeline: !GenericPipeline\n  runs:\n    - *prepare_stage # prepare\n    - *train_stage\n    - *present_model # finals\n",
		GitPath:   "https://gitlab.com/mlrep/template-default/-/blob/master/experiment.yml",
		CreatedAt: timestamppb.New(time.Date(2023, 11, 10, 12, 23, 11, 12, time.UTC)),
		UpdatedAt: timestamppb.New(time.Now()),
	}, nil
}

func (g *IntegratorHandler) GetHello(context.Context, *pbintegrator.GetHelloRequest) (*pbintegrator.GetHelloResponse, error) {
	return &pbintegrator.GetHelloResponse{Pong: "fek"}, nil
}

func GenerateRandomHash() string {
	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создаем случайную строку (например, 20 байт случайных данных)
	randomData := make([]byte, 20)
	for i := range randomData {
		randomData[i] = byte(rand.Intn(256))
	}

	// Хэшируем данные с использованием SHA-1
	hasher := sha1.New()
	hasher.Write(randomData)
	hash := hasher.Sum(nil)

	// Возвращаем хэш в виде строки (hex)
	return hex.EncodeToString(hash)
}
