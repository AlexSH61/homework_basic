package main

import (
	"fmt"

	"github.com/AlexSH61/homework_basic/hw02_fix_app/printer"
	"github.com/AlexSH61/homework_basic/hw02_fix_app/reader"
)

func main() {
	var path string = "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)
	if len(path) == 0 {
		path = "data.json"
	}

	staff, err := reader.ReadJSON(path)
	if err != nil {
		fmt.Print(err)
	}
	printer.PrintStaff(staff)
}
