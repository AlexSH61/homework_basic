package comparator

import "github.com/AlexSH61/homework_basic/book"

type ByWhat int

const (
	ByYear ByWhat = iota
	BySize
	ByRating
)

type BookComparison struct {
	mode ByWhat
}

func NewCompare(mode ByWhat) *BookComparison {
	return &BookComparison{mode}
}

func (bc *BookComparison) Compare(firstBook, secondBook book.Book) bool {
	switch bc.mode {
	case ByYear:
		return firstBook.GetYear() > secondBook.GetYear()
	case BySize:
		return firstBook.GetSize() > secondBook.GetSize()
	case ByRating:
		return firstBook.GetRating() > secondBook.GetRating()
	default:
		return false
	}
}
