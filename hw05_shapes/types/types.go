package types

import (
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
func NewCircle(radius float64) *Circle {
	return &Circle{
		radius: radius,
	}
}

type Triangle struct {
	base   float64
	height float64
}

func NewTriangle(base, height float64) *Triangle {
	return &Triangle{
		base:   base,
		height: height,
	}
}
func (t *Triangle) Area() float64 {
	return t.base * t.height / 2
}

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width, heigth float64) *Rectangle {
	return &Rectangle{
		width:  width,
		height: heigth,
	}
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}
