package repository

import (
	"context"
	"go-advanced-shapes/pkg/models"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/mock"
)

type MockDynamoDB struct {
	mock.Mock
}

func (m *MockDynamoDB) ListShapesByType(shapeType string) ([]models.Request, error) {
	args := m.Called(shapeType)
	return args.Get(0).([]models.Request), args.Error(1)
}

func (m *MockDynamoDB) CreateShape(id string, shapeType string, a float64, b float64, creator string) error {
	args := m.Called(id, shapeType, a, b, creator)
	return args.Error(0)
}

type MockClient struct {
	mock.Mock
}

func (m *MockClient) Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	args := m.Called(ctx, params, optFns)
	return args.Get(0).(*dynamodb.QueryOutput), args.Error(1)
}

func (m *MockClient) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	args := m.Called(ctx, params, optFns)
	return args.Get(0).(*dynamodb.PutItemOutput), args.Error(1)
}

func (m *MockClient) ExecuteStatement(ctx context.Context, params *dynamodb.ExecuteStatementInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ExecuteStatementOutput, error) {
	args := m.Called(ctx, params, optFns)
	return args.Get(0).(*dynamodb.ExecuteStatementOutput), args.Error(1)
}
