package repository

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/mock"
)

type MockS3Repository struct {
	mock.Mock
}

func (m *MockS3Repository) UploadFile(bucketName string, objectKey string, fileReader io.Reader) error {
	args := m.Called(bucketName, objectKey, fileReader)
	return args.Error(0)
}

type MockClientS3 struct {
	mock.Mock
}

func (m *MockClientS3) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	args := m.Called(ctx, params, optFns)
	return args.Get(0).(*s3.PutObjectOutput), args.Error(1)
}
