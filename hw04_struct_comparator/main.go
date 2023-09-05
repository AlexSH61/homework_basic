package main

import (
	"fmt"

	"github.com/AlexSH61/homework_basic/book"
	comparator "github.com/AlexSH61/homework_basic/p_const"
)

func main() {
	// firstBook := Book{1, "War and Peace", "Leo Tolstoy", 1869, 1225, 4.2}
	// secondBook := Book{2, "Crime and Punishment", "Fyodor Dostoevsky", 1866, 671, 4.5}

	FirstBook := book.NewBook(1, "War and Peace", "Leo Tolstoy", 1869, 1225, 4.2)
	SecondBook := book.NewBook(2, "Crime and Punishment", "Fyodor Dostoevsky", 1866, 671, 4.5)
	fmt.Println(FirstBook.GetAuthor())
	fmt.Println(SecondBook)
	bc := comparator.BookComparison{Mode: (comparator.ByRating)}
	fmt.Println(bc.Compare(*FirstBook, *SecondBook))

}
