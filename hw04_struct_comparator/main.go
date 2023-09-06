package main

import (
	"fmt"

	"github.com/AlexSH61/homework_basic/book"
	comparator "github.com/AlexSH61/homework_basic/p_const"
)

func main() {
	firstBook := book.NewBook(1, "War and Peace", "Leo Tolstoy", 1869, 1225, 4.2)
	secondBook := book.NewBook(2, "Crime and Punishment", "Fyodor Dostoevsky", 1866, 671, 4.5)
	firstBook.Rating(4.3)
	secondBook.Rating(5)
	fmt.Println(firstBook.GetAuthor())
	firstBook.Author("Lev Tolstoy")
	fmt.Println(firstBook.GetAuthor())
	fmt.Println(secondBook)
	bc := comparator.NewCompare(comparator.ByRating)
	fmt.Println(bc.Compare(*firstBook, *secondBook))
	fmt.Println(firstBook.GetRating())
}
