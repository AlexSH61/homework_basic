package book_test

import (
	"testing"

	"github.com/AlexSH61/homework_basic/hw06_testing/hw04_testing/type_book"
	"github.com/stretchr/testify/assert"
)

func TestBookType(t *testing.T) {
	book := book.NewBook(3, "Master and Margarita", "Michail Bulgakov", 1867, 500, 4.5)

	book.ID(3)
	assert.Equal(t, 3, book.GetID())

	book.Title("Master and Margarita")
	assert.Equal(t, "Master and Margarita", book.GetTitle())

	book.Author("Mikhail Bulgakov")
	assert.Equal(t, "Mikhail Bulgakov", book.GetAuthor())

	book.Year(1867)
	assert.Equal(t, 1867, book.GetYear())

	book.Size(500)
	assert.Equal(t, 500, book.GetSize())

	book.Rating(4.5)
	assert.Equal(t, float64(4.5), book.GetRating())
}
