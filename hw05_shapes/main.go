package main

import (
	"fmt"

	printShape "github.com/AlexSH61/homework_basic-1/calculateArea"
	"github.com/AlexSH61/homework_basic-1/types"
)

func main() {
	c := types.NewCircle(5)
	r := types.NewRectangle(8, 5)
	t := types.NewTriangle(6, 4)
	sliceArea := []types.Shape{c, r, t}
	for _, area := range sliceArea {
		if sliceArea, err := printShape.СalculateArea(area); err != nil {
			fmt.Println("incorrect type", err)
		} else {
			fmt.Printf("area Circle:%T %v\n", area, sliceArea)
		}
	}
}
