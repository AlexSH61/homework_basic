package jsonserializer

import (
	"testing"

	"github.com/AlexSH61/homework_basic/hw09_serialize/book"
	"github.com/stretchr/testify/assert"
)

func TestJsonSerDes(t *testing.T) {
	testBooks := []book.Book{
		{
			ID:     1,
			Title:  "The Go Programming Language",
			Author: "Alan A. A. Donovan",
			Year:   2015,
			Size:   356,
			Rate:   4.5,
		},
		{
			ID:     2,
			Title:  "Clean Code: A Handbook of Agile Software Craftsmanship",
			Author: "Robert C. Martin",
			Year:   2008,
			Size:   464,
			Rate:   4.8,
		},
	}

	serializeData, err := SerializeBookToJSON(testBooks)
	assert.NoError(t, err, "ошибка при сериализации")
	resultBooks1, err := DeserializeJSONToBook(serializeData)
	assert.NoError(t, err, "ошибка при десериализации")
	assert.Equal(t, testBooks, resultBooks1)
}
