package main

import (
	"go-advanced-shapes/internal/processor"

	d "go-advanced-shapes/internal/repository/dynamodb"
	s "go-advanced-shapes/internal/repository/s3"
	"go-advanced-shapes/pkg/handler"
	"go-advanced-shapes/pkg/services"
	"go-advanced-shapes/pkg/validator"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	dy := d.NewDynamoDB()
	s3r := s.NewS3Repository()
	us := services.NewUserDataService(&http.Client{})
	p := processor.New(dy, s3r, us)
	v := validator.NewItemValidator()
	handler := handler.New(p, v)

	lambda.Start(handler.HandleApiRest)
}
