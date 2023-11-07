package models

import (
	"fmt"
	"strings"
)

func NewEllipse(id string, a float64, b float64) IShape {
	return &Ellipse{
		Shape{id},
		a,
		b,
	}
}

func NewRectangle(id string, a float64, b float64) IShape {
	return &Rectangle{
		Shape{id},
		a,
		b,
	}
}

func NewTriangle(id string, a float64, b float64) IShape {
	return &Triangle{
		Shape{id},
		a,
		b,
	}
}

func ShapeFactory(id string, shapeType string, a float64, b float64) (IShape, error) {
	shapeTypelower := strings.ToLower(shapeType)
	if shapeTypelower == "ellipse" {
		return NewEllipse(id, a, b), nil
	}
	if shapeTypelower == "rectangle" {
		return NewRectangle(id, a, b), nil
	}
	if shapeTypelower == "triangle" {
		return NewTriangle(id, a, b), nil
	}
	return nil, fmt.Errorf("ShapeFactory. Invalid shape type (%s).", shapeType)
}
