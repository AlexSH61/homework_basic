package main

import (
	"fmt"

	printArea "github.com/AlexSH61/homework_basic/calculateArea"
	"github.com/AlexSH61/homework_basic/types"
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
	areaCircle, err := printArea.CalculateAre(circle)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Printf("%T: ", circle)
		fmt.Println("Area:", areaCircle)
	}
	areaRectangle, err := printArea.CalculateAre(rectangle)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Printf("%T: ", rectangle)
		fmt.Println("Area:", areaRectangle)
	}
	areaTringle, err := printArea.CalculateAre(triangle)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Printf("%T: ", triangle)
		fmt.Println("Area:", areaTringle)
	}
}
