package main

import (
	"fmt"

	binaryAlg "github.com/AlexSH61/homework_basic/binary_search/binary"
)

func main() {
	result := binaryAlg.BinarySearching(5, []int{2, 3, 4, 5, 6, 7, 8, 89})
	fmt.Println(result)
}
