package jsonserializer_test

import (
	"os"
	"testing"

	"github.com/AlexSH61/homework_basic/hw09_serialize/book"
	"github.com/AlexSH61/homework_basic/hw09_serialize/jsonserializer"
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
	jsonFilePath := "test_books.json"
	defer func() {
		if err := os.Remove(jsonFilePath); err != nil {
			t.Error("ошибка при удаление ", err)
		}
	}()
	jsonserializer.SerializeBookToJSON(testBooks, jsonFilePath)
	_, err := os.Stat(jsonFilePath)
	assert.NoError(t, err, "ошибка при создание файла")
	resultBooks1, err := jsonserializer.DeserializeJSONToBook(jsonFilePath)
	assert.NoError(t, err, "ошибка при сериализации/десериализации")
	assert.Equal(t, testBooks, resultBooks1)
}
