package protobufserializer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AlexSH61/homework_basic/hw09_serialize/bookpb"
)

func TestJsonSerDes(t *testing.T) {
	testBooks := &bookpb.Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan",
	}
	serBook, err := SerializeBookToProtobuf(testBooks)
	assert.NoError(t, err)
	desBook, err := DeserializeProtobufToBook(serBook)
	assert.NoError(t, err)
	assert.Equal(t, testBooks.Title, desBook.Title)
	assert.Equal(t, testBooks.Author, desBook.Author)
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
	serProtoSliceBooks, err := SerializeBookSliceToProtobuf(testBooksArray)
	assert.NoError(t, err)
	desProtoSliceBooks, err := DeserializeProtobufSliceToBook(serProtoSliceBooks)
	assert.NoError(t, err)
	assert.Len(t, desProtoSliceBooks, len(testBooksArray))
	for i := range testBooksArray {
		assert.Equal(t, testBooksArray[i].Author, desProtoSliceBooks[i].Author)
		assert.Equal(t, testBooksArray[i].Title, desProtoSliceBooks[i].Title)
	}
}
