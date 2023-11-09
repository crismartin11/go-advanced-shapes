package repository

import (
	"go-advanced-shapes/pkg/models"
	service "go-advanced-shapes/pkg/services/serviceForAdapter"

	"github.com/rs/zerolog/log"
)

type AdapterDynamoDB struct {
	client Client
}

func NewAdapterDynamoDB() IDynamoDB {
	return AdapterDynamoDB{}
}

func (d AdapterDynamoDB) ListShapesByType(shapeType string) ([]models.Request, error) {
	log.Info().Msg("ListShapesByType -- PartiQl --")

	// Note: this method should do some transform because it's an adapter

	s := service.NewServiceForAdapter()
	return s.ListShapesByTypePartiQl(shapeType)
}

func (d AdapterDynamoDB) CreateShape(id string, shapeType string, a float64, b float64, creator string) error {
	log.Info().Msg("CreateShape -- PartiQl --")
	s := service.NewServiceForAdapter()

	// Note: this method should do some transform because it's an adapter
	//TODO: armar y pasar objeto

	return s.CreateShapePartiQl(id, shapeType, a, b, creator)
}
