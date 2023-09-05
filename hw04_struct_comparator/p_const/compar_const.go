package comparator

import "github.com/AlexSH61/homework_basic/book"

type ByWhat int

// enum.
const (
	ByYear ByWhat = iota
	BySize
	ByRating
)

// create const struct.
type BookComparison struct {
	Mode ByWhat
}

// create const constructor.
func NewCompare(Mode ByWhat) *BookComparison {
	return &BookComparison{Mode}
}

// creat method сomprasion.
func (bc *BookComparison) Compare(FirstBook, SecondBook book.Book) bool {
	switch bc.Mode {
	case ByYear:
		return book.Book.GetYear(FirstBook) > book.Book.GetYear(SecondBook)
	case BySize:
		return book.Book.GetSize(FirstBook) > book.Book.GetSize(SecondBook)
	case ByRating:
		return book.Book.GetRating(FirstBook) > book.Book.GetRating(SecondBook)
	default:
		// panic("incorrect mode").
		return false
	}
}
