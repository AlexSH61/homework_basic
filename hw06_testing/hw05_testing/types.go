package types

import (
	"errors"
	"math"
)

var (
	ErrIncorectRadius = errors.New("incorrect radius")
	ErrIncorectSide   = errors.New("incorrect side")
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func NewCircle(radius float64) (*Circle, error) {
	if radius <= 0 {
		return nil, ErrIncorectRadius
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
		return nil, ErrIncorectSide
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

func NewRectangle(width, height float64) (*Rectangle, error) {
	if height <= 0 || width <= 0 {
		return nil, ErrIncorectSide
	}
	return &Rectangle{
		width:  width,
		height: height,
	}, nil
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}
