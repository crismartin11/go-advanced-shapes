package models

import (
	"fmt"
	"math"
)

type IShape interface {
	Area() float64
	Detail() string
	DetailPrint()
}

type Shape struct {
	ID string
}

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

func (e Ellipse) DetailPrint() {
	fmt.Println("Ellipse")
	fmt.Printf("- ID: %s\n- Radio A: %.2f\n- Radio B: %.2f\n- Area: %.2f\n", e.ID, e.RadioA, e.RadioB, e.Area())
}

type Rectangle struct {
	Shape
	High float64
	Long float64
}

func (r Rectangle) Area() float64 {
	return r.High * r.High
}

func (r Rectangle) Detail() string {
	return fmt.Sprintf("- Rectangle - ID: %s - High: %.2f - Long: %.2f - Area: %.2f", r.ID, r.High, r.Long, r.Area())
}

func (r Rectangle) DetailPrint() {
	fmt.Println("Rectangle")
	fmt.Printf("- ID: %s\n- High: %.2f\n- Long: %.2f\n- Area: %.2f\n", r.ID, r.High, r.Long, r.Area())
}

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

func (t Triangle) DetailPrint() {
	fmt.Println("Triangle")
	fmt.Printf("- ID: %s\n- Base: %.2f\n- Height: %.2f\n- Area: %.2f\n", t.ID, t.Base, t.Height, t.Area())
}
