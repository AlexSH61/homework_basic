package main

import (
	"fmt"

	"github.com/AlexSH61/homework_basic/book"
	comparator "github.com/AlexSH61/homework_basic/p_const"
)

func main() {
	// firstBook := Book{1, "War and Peace", "Leo Tolstoy", 1869, 1225, 4.2}
	// secondBook := Book{2, "Crime and Punishment", "Fyodor Dostoevsky", 1866, 671, 4.5}
	firstBook := book.NewBook(1, "War and Peace", "Leo Tolstoy", 1869, 1225, 4.2)
	secondBook := book.NewBook(2, "Crime and Punishment", "Fyodor Dostoevsky", 1866, 671, 4.5)
	firstBook.Rating(4.9)
	fmt.Println(firstBook.GetAuthor())
	fmt.Println(secondBook)
	bc := comparator.NewCompare(comparator.ByRating)
	fmt.Println(bc.Compare(*firstBook, *secondBook))
	fmt.Println(firstBook.GetRating())
}
