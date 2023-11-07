package processor

import (
	"context"
	"go-advanced-shapes/pkg/models"

	"github.com/stretchr/testify/mock"
)

type MockProcessor struct {
	mock.Mock
}

func (m *MockProcessor) ProcessCreation(request models.Request) (models.Response, error) {
	args := m.Called(request)
	return args.Get(0).(models.Response), args.Error(1)
}

func (m *MockProcessor) ProcessGeneration(ctx context.Context, request models.Request) (models.Response, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(models.Response), args.Error(1)
}
