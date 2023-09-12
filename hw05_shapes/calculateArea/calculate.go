package printArea

import (
	"fmt"

	"github.com/AlexSH61/homework_basic/types"
)

func CalculateAre(shape types.Shape) (float64, error) {
	if shape, ok := shape.(types.Shape); ok {
		return shape.Area(), nil
	}
	return 0, fmt.Errorf("incorrect shape")
}
