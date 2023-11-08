package service

import (
	"fmt"
	"go-advanced-shapes/pkg/models"
)

type ServiceForAdapter struct{}

type IServiceForAdapter interface {
	ListShapesByTypePartiQl(shapeType string) ([]models.Request, error)
	CreateShapePartiQl(id string, shapeType string, a float64, b float64, creator string) error // TODO: pasar objeto
}

func NewServiceForAdapter() IServiceForAdapter {
	return ServiceForAdapter{}
}

func (s ServiceForAdapter) ListShapesByTypePartiQl(shapeType string) ([]models.Request, error) {
	fmt.Println("ListShapesByType del service for adapter")
	return []models.Request{}, nil
}

func (s ServiceForAdapter) CreateShapePartiQl(id string, shapeType string, a float64, b float64, creator string) error {
	fmt.Println("CreateShape del service for adapter")
	return nil
}
