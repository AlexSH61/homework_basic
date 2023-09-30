package types_test

import (
	"errors"
	"testing"

	types "github.com/AlexSH61/homework_basic/hw06_testing/hw05_testing"
)

func TestCircle(t *testing.T) {
	checkCircle := []struct {
		name   string
		radius float64
		want   float64
		err    error
	}{
		{name: "Circle", radius: 5.0, want: 78.53981633974483, err: nil},
		{name: "incorrectCircle", radius: 0.0, want: 0.0, err: types.ErrIncorectRadius},
		{name: "incorrectCircle", radius: -1.0, want: 0.0, err: types.ErrIncorectRadius},
	}
	for _, cc := range checkCircle {
		t.Run(cc.name, func(t *testing.T) {
			c, err := types.NewCircle(cc.radius)

			if errors.Is(err, types.ErrIncorectRadius) {
				return
			}
			if cc.err == nil {
				hasArea := c.Area()

				if hasArea != cc.want {
					t.Errorf("area %f, but got %f", cc.want, hasArea)
				}
			}
		})
	}
}

func TestTriangle(t *testing.T) {
	checkTriangle := []struct {
		name   string
		base   float64
		height float64
		want   float64
		err    error
	}{
		{name: "Triangle", base: float64(5), height: float64(10), want: 25, err: nil},
		{name: "incorrectTrinagle", base: float64(-1), height: float64(10), want: 0.0, err: types.ErrIncorectSide},
		{name: "incorrectTriangle", base: float64(0), height: float64(10), want: 0.0, err: types.ErrIncorectSide},
	}

	for _, ct := range checkTriangle {
		t.Run(ct.name, func(t *testing.T) {
			c, err := types.NewTriangle(ct.base, ct.height)

			if errors.Is(err, types.ErrIncorectSide) {
				return
			}

			if ct.err == nil {
				hasArea := c.Area()

				if hasArea != ct.want {
					t.Errorf("area %f, but got %f", ct.want, hasArea)
				}
			}
		})
	}
}

func TestRectangleArea(t *testing.T) {
	checkRectangle := []struct {
		name   string
		width  float64
		height float64
		want   float64
		err    error
	}{
		{name: "Rectangle", width: 5.0, height: 10.0, want: 50.0, err: nil},
		{name: "IncorrectRectangle", width: -1.0, height: 4.0, want: 0.0, err: types.ErrIncorectSide},
		{name: "IncorrectRectangle", width: 0.0, height: 0.0, want: 0.0, err: types.ErrIncorectSide},
	}

	for _, cr := range checkRectangle {
		t.Run(cr.name, func(t *testing.T) {
			r, err := types.NewRectangle(cr.width, cr.height)
			if errors.Is(err, types.ErrIncorectSide) {
				return
			}

			if cr.err == nil {
				hasArea := r.Area()

				if hasArea != cr.want {
					t.Errorf("area %f, but got %f", cr.want, hasArea)
				}
			}
		})
	}
}
