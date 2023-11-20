package protobufserializer_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AlexSH61/homework_basic/hw09_serialize/bookpb"
	"github.com/AlexSH61/homework_basic/hw09_serialize/protobufserializer"
)

func TestJsonSerDes(t *testing.T) {
	testBooks := &bookpb.Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan",
	}
	testFile := "test_books.bin"
	defer func() {
		_ = os.Remove(testFile)
	}()
	err := protobufserializer.SerializeBookToProtobuf(testBooks, testFile)
	assert.NoError(t, err)

	assert.FileExists(t, testFile)
	deserializedBook, err := protobufserializer.DeserializeProtobufToBook(testFile)
	assert.NoError(t, err)
	assert.Equal(t, testBooks.Title, deserializedBook.Title)
	assert.Equal(t, testBooks.Author, deserializedBook.Author)
}
func TestProtoSliceSerDes(t *testing.T) {
	testBooksArray := []*bookpb.Book{
		{
			Title:  "Clean Code: A Handbook of Agile Software Craftsmanship",
			Author: "Robert C. Martin",
		}, {
			Title:  "The Go Programming Language",
			Author: "Alan A. A. Donovan",
		},
	}
	testFile := "testprotoSLice.bin"
	defer func() {
		_ = os.Remove(testFile)
	}()
	err := protobufserializer.SerializeBookSliceToProtobuf(testBooksArray, testFile)
	assert.NoError(t, err)
	assert.FileExists(t, testFile)
	deserializeSliceBooks, err := protobufserializer.DeserializeProtobufSliceToBook(testFile)
	assert.NoError(t, err)
	assert.Len(t, deserializeSliceBooks, len(testBooksArray))
	for i := range testBooksArray {
		assert.Equal(t, testBooksArray[i].Author, deserializeSliceBooks[i].Author)
		assert.Equal(t, testBooksArray[i].Title, deserializeSliceBooks[i].Title)
	}
}
