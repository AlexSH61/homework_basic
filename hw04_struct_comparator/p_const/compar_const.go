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
	mode ByWhat
}

// create const constructor.
func NewCompare(mode ByWhat) *BookComparison {
	return &BookComparison{mode}
}

// creat method сomprasion.
func (bc *BookComparison) Compare(FirstBook, SecondBook book.Book) bool {
	switch bc.mode {
	case ByYear:
		return FirstBook.GetYear() > SecondBook.GetYear()
	case BySize:
		return FirstBook.GetSize() > SecondBook.GetSize()
	case ByRating:
		return FirstBook.GetRating() > FirstBook.GetRating()
	default:
		return false
	}
}
