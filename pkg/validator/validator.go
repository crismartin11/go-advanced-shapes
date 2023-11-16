package validator

import (
	"fmt"
	"go-advanced-shapes/pkg/models"

	"github.com/rs/zerolog/log"
)

type Validator interface {
	ValidateRequest(request models.Request) error
}

type ItemValidator struct{}

func NewItemValidator() Validator {
	return &ItemValidator{}
}

func (i ItemValidator) ValidateRequest(request models.Request) error {
	sublogger := log.With().Str("component", "ValidateRequest").Logger()
	sublogger.Info().Msg("start")

	if !request.IsValidShapeType() {
		sublogger.Error().Str("ShapeType", request.ShapeType).Msg("Handle Shape. Invalid shape type.")
		return fmt.Errorf("ERROR: Tipo de figura %s inválido.", request.ShapeType)
	}
	return nil
}
