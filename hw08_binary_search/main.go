package main

import (
	"fmt"

	binary "github.com/AlexSH61/homework_basic/hw08_binary_search/binarySearch"
)

func main() {
	result := binary.Searching(5, []int{2, 3, 4, 5, 6, 7, 8, 89})
	fmt.Println(result)
}
