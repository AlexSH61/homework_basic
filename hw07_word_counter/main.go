package main

import (
	"fmt"

	counter "github.com/AlexSH61/homework_basic/hw07_word_counter/word_counter"
)

const text = "The night is darkest just before the dawn. And I promise you, the dawn is coming."

func main() {
	fmt.Println(counter.CountWords(text))
}
