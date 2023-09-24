package compare_test

import (
	"testing"

	"github.com/AlexSH61/homework_basic/book"

	comparator "github.com/AlexSH61/homework_basic/hw06_testing/compare"
)

func TestComparator(t *testing.T) {
	tests := []struct {
		name       string
		firstBook  book.Book
		secondBook book.Book
		mode       comparator.ByWhat
		answer     bool
	}{
		{
			name:       "Year in firstBook > Year in secondBook",
			firstBook:  *book.NewBook(1, "War and Peace", "Leo Tolstoy", 1869, 1225, 4.2),
			secondBook: *book.NewBook(2, "Crime and Punishment", "Fyodor Dostoevsky", 1866, 671, 4.5),
			mode:       comparator.ByYear,
			answer:     true,
		},
		{
			name:       "Rating in firstBook > Rating in secondBook",
			firstBook:  *book.NewBook(1, "War and Peace", "Leo Tolstoy", 1869, 1225, 4.1),
			secondBook: *book.NewBook(2, "Crime and Punishment", "Fyodor Dostoevsky", 1866, 671, 4.5),
			mode:       comparator.ByRating,
			answer:     false,
		},
		{
			name:       "Size in firstBook > Size in secondBook",
			firstBook:  *book.NewBook(1, "War and Peace", "Leo Tolstoy", 1869, 1225, 4.2),
			secondBook: *book.NewBook(2, "Crime and Punishment", "Fyodor Dostoevsky", 1866, 671, 4.5),
			mode:       comparator.BySize,
			answer:     true,
		},
		{
			name:       "If Rating in firstBook = Rating in secondBook",
			firstBook:  *book.NewBook(1, "Name Sample", "Ivanov Ivan", 1869, 1225, 4.6),
			secondBook: *book.NewBook(1, "Sample Name", "Petrov Petr", 1866, 656, 4.6),
			mode:       comparator.ByRating,
			answer:     false,
		},
	}

	for _, spl := range tests {
		t.Run(spl.name, func(t *testing.T) {
			comparate := comparator.NewCompare(spl.mode)
			result := comparate.Compare(spl.firstBook, spl.secondBook)
			if result != spl.answer {
				t.Errorf("%s: want %v, but got %v", spl.name, spl.answer, result)
			}
		})
	}
}
