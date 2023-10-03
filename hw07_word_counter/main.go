package main

import (
	"fmt"
	"math/rand"

	prepare "github.com/AlexSH61/homework_basic/hw07_word_counter/prepare_string"
	counter "github.com/AlexSH61/homework_basic/hw07_word_counter/word_counter"
)

func getRandText() string {
	texts := []string{
		"Everything will be alright in the end. If it's the end.",
		"The night is darkest just before the dawn. And I promise you, the dawn is coming.",
		"",
	}
	randomIndex := rand.Intn(3)
	return texts[randomIndex]
}

func main() {
	preparedText, err := prepare.StringPrepare(getRandText())
	if err != nil {
		fmt.Println("Error in PrepareString:", err)
		return
	}
	fmt.Println(counter.CountWords(preparedText))
}
