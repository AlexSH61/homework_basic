package main

import (
	"fmt"
	"os"

	"github.com/AlexSH61/homework_basic/hw03_chessboard/board"
)

func main() {
	fmt.Println("Enter number:")
	var size int
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Println("incorrect input, please try again")
	}
	defer os.Exit(size)
	fmt.Println(board.Boardsize(size))
}
