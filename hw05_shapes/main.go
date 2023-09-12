package main

import (
	"fmt"

	calculatearea "github.com/AlexSH61/homework_basic/hw05_shapes/calculateArea"
	"github.com/AlexSH61/homework_basic/hw05_shapes/types"
)

func main() {
	circle, err := types.NewCircle(5)
	if err != nil {
		fmt.Println("Error", err)
	}
	rectangle, err := types.NewRactangle(10, 5)
	if err != nil {
		fmt.Println("Error", err)
	}
	triangle, err := types.NewTriangle(10, 5)
	if err != nil {
		fmt.Println("Error", err)
	}
	areaCircle, err := calculatearea.CalculateAre(circle)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Printf("%T: ", circle)
		fmt.Println("Area:", areaCircle)
	}
	areaRectangle, err := calculatearea.CalculateAre(rectangle)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Printf("%T: ", rectangle)
		fmt.Println("Area:", areaRectangle)
	}
	areaTringle, err := calculatearea.CalculateAre(triangle)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Printf("%T: ", triangle)
		fmt.Println("Area:", areaTringle)
	}
}
