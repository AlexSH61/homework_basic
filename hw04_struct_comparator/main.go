package main

import (
	"fmt"
)

// creat book struct.
type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rating float32
}

// create constructor book struct.
func NewBook(id int, title, author string, year, size int, rating float32) Book {
	return Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rating: rating,
	}
}

// creat custom type.
type byWhat int

// enum.
const (
	byYear byWhat = iota
	bySize
	byRating
)

// create const struct.
type BookComparison struct {
	mode byWhat
}

// create const constructor.
func NewCompare(mode byWhat) BookComparison {
	return BookComparison{mode}
}

// creat method сomprasion.
func (bc BookComparison) Compare(firstBook, secondBook Book) bool {
	switch bc.mode {
	case byYear:
		return firstBook.year > secondBook.year
	case bySize:
		return firstBook.size > secondBook.size
	case byRating:
		return firstBook.rating > secondBook.rating
	default:
		// panic("incorrect mode").
		return false
	}
}

// create set&get feilds struct.
func (b *Book) SetID(id int) {
	b.id = id
}

func (b Book) GetID() int {
	return b.id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b Book) GetTitle() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b Book) GetAuthor() string {
	return b.author
}

func (b *Book) SetSize(year int) {
	b.year = year
}

func (b Book) GetSize() int {
	return b.size
}

func (b *Book) SetRate(rating float32) {
	b.rating = rating
}

func (b Book) GetRating() float32 {
	return b.rating
}

func main() {
	firstBook := Book{1, "War and Peace", "Leo Tolstoy", 1869, 1225, 4.2}
	secondBook := Book{2, "Crime and Punishment", "Fyodor Dostoevsky", 1866, 671, 4.5}
	fmt.Println(firstBook.author)
	fmt.Println(secondBook)
	bc := BookComparison{mode: byRating}
	fmt.Println(bc.Compare(firstBook, secondBook))
}
