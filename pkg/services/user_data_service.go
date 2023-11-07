package services

import (
	"encoding/json"
	"fmt"
	"go-advanced-shapes/pkg/constants"
	"go-advanced-shapes/pkg/models"
	"io"
	"net/http"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type UserDataService struct {
	h HttpClient
}

type IUserDataService interface {
	GetUserData(id string) (models.UserDataResponse, error)
}

func NewUserDataService(client HttpClient) IUserDataService {
	return UserDataService{
		h: client,
	}
}

func (us UserDataService) GetUserData(id string) (models.UserDataResponse, error) {

	var userData models.UserDataResponse

	url := constants.URL_API_USERS + id

	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	httpResponse, err := us.h.Do(request)
	if err != nil {
		return userData, fmt.Errorf("GetUserData. Error consultando datos de api (%s). %s", url, err)
	}

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return userData, fmt.Errorf("GetUserData. Error leyendo datos de response. %s", err)
	}

	err = json.Unmarshal(body, &userData)
	if err != nil {
		return userData, fmt.Errorf("GetUserData. Error parseando datos de response. %s", err)
	}

	return userData, nil
}
