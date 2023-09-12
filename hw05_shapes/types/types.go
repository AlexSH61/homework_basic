package types

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	radius float64
}

func NewCircle(radius float64) (*Circle, error) {
	if radius <= 0 {
		return nil, fmt.Errorf("incorrect radradius")
	}
	return &Circle{
		radius: radius,
	}, nil
}
func (c *Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

type Triangle struct {
	base   float64
	height float64
}

func NewTriangle(base, height float64) (*Triangle, error) {
	if base <= 0 || height <= 0 {
		return nil, fmt.Errorf("incorrect radradius")
	}
	return &Triangle{
		base:   base,
		height: height,
	}, nil
}
func (t *Triangle) Area() float64 {
	return t.height * t.base / 2
}

type Rectangle struct {
	width  float64
	height float64
}

func NewRactangle(width, height float64) (*Rectangle, error) {
	if height <= 0 || width <= 0 {
		return nil, fmt.Errorf("incorrect radradius")
	}
	return &Rectangle{
		width:  width,
		height: height,
	}, nil
}
func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

// type Square struct {
// 	side float64
// }
// func NewSquare(side float64) *Square{
// 	return &Square{
// 		side: side,
// 	}
// }
// func (s Square)  float64 {
// 	return s.side * s.side
// }
