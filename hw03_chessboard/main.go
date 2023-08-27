package main

import (
	"fmt"

	"github.com/AlexSH61/homework_basic/hw03_chessboard/board"
)

func main() {
	fmt.Println("Enter number:")
	var size int
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		panic(err)
	}
	fmt.Println(board.BoardSize(size))
	// Place your code here.
}
