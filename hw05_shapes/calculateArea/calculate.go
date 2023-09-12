package calculatearea

import (
	"fmt"

	"github.com/AlexSH61/homework_basic/hw05_shapes/types"
)

func CalculateAre(shape any) (float64, error) {
	if shape, ok := shape.(types.Shape); ok {
		return shape.Area(), nil
	}
	return 0, fmt.Errorf("incorrect shape")
}
