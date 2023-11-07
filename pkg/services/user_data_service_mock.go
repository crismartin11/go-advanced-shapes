package services

import (
	"go-advanced-shapes/pkg/models"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockUserDataService struct {
	mock.Mock
}

func (m *MockUserDataService) GetUserData(id string) (models.UserDataResponse, error) {
	args := m.Called(id)
	return args.Get(0).(models.UserDataResponse), args.Error(1)
}

type MockHttpClient struct {
	mock.Mock
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}
