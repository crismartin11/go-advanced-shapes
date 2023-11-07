package models

import "fmt"

type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func NewResponseOk(body string) (Response, error) {
	return Response{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       fmt.Sprintf(body),
	}, fmt.Errorf("")
}

func NewResponseError(statusCode int, err string) (Response, error) {
	return Response{
		StatusCode: statusCode,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       "",
	}, fmt.Errorf(err)
}
