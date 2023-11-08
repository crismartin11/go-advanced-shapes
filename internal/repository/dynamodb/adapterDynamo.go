package repository

import (
	"context"
	"go-advanced-shapes/pkg/models"
	service "go-advanced-shapes/pkg/services/serviceForAdapter"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/rs/zerolog/log"
)

type AdapterDynamoDB struct {
	client Client
}

func NewAdapterDynamoDB() IDynamoDB {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("Configuration error, " + err.Error())
	}

	return AdapterDynamoDB{client: dynamodb.NewFromConfig(cfg)}
}

func (d AdapterDynamoDB) ListShapesByType(shapeType string) ([]models.Request, error) {

	log.Info().Msg("ListShapesByType -- PartiQl --")
	s := service.NewServiceForAdapter()
	return s.ListShapesByTypePartiQl(shapeType)
}

func (d AdapterDynamoDB) CreateShape(id string, shapeType string, a float64, b float64, creator string) error {

	log.Info().Msg("CreateShape -- PartiQl --")
	s := service.NewServiceForAdapter()
	//TODO: armar y pasar objeto

	return s.CreateShapePartiQl(id, shapeType, a, b, creator)
}
