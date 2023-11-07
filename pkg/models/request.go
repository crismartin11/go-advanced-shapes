package models

import (
	"strings"

	"github.com/rs/zerolog/log"
)

const ELLIPSE string = "ELLIPSE"
const RECTANGLE string = "RECTANGLE"
const TRIANGLE string = "TRIANGLE"

type Request struct {
	Id        string  `json:"id" dynamodbav:"id"`
	ShapeType string  `json:"tipo" dynamodbav:"tipo"`
	A         float64 `json:"a" dynamodbav:"a"`
	B         float64 `json:"b" dynamodbav:"b"`
	Creator   string  `json:"creador" dynamodbav:"creador"`
}

func (r Request) IsValidShapeType() bool {
	shapeType := strings.ToUpper(r.ShapeType)
	log.Error().Str("ShapeType Upper", shapeType).Msg("Handle Shape. Invalid shape type.")
	return shapeType == ELLIPSE || shapeType == RECTANGLE || shapeType == TRIANGLE
}

func (r Request) IsValidData() bool {
	return r.Id != "" && r.A != 0 && r.B != 0
}
