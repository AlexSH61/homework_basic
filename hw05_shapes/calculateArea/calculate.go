package printShape

import (
	"fmt"

	"github.com/AlexSH61/homework_basic-1/types"
)

func СalculateArea(shape types.Shape) (float64, error) {
	if shape == nil {
		return 0, fmt.Errorf("unknown shape")
	}
	return shape.Area(), nil
}
