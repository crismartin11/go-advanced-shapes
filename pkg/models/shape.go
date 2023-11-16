package models

import (
	"fmt"
	"math"
)

type IShape interface {
	Area() float64
	Detail() string
}

type Shape struct {
	ID string
}

// Ellipse
type Ellipse struct {
	Shape
	RadioA float64
	RadioB float64
}

func (e Ellipse) Area() float64 {
	pi := math.Pi
	return pi * e.RadioA * e.RadioB
}

func (e Ellipse) Detail() string {
	return fmt.Sprintf("- Ellipse - ID: %s - Radio A: %.2f - Radio B: %.2f - Area: %.2f", e.ID, e.RadioA, e.RadioB, e.Area())
}

// Rectangle
type Rectangle struct {
	Shape
	High float64
	Long float64
}

func (r Rectangle) Area() float64 {
	return r.High * r.Long
}

func (r Rectangle) Detail() string {
	return fmt.Sprintf("- Rectangle - ID: %s - High: %.2f - Long: %.2f - Area: %.2f", r.ID, r.High, r.Long, r.Area())
}

// Triangle
type Triangle struct {
	Shape
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) Detail() string {
	return fmt.Sprintf("- Triangle - ID: %s - Base: %.2f - Height: %.2f - Area: %.2f", t.ID, t.Base, t.Height, t.Area())
}
