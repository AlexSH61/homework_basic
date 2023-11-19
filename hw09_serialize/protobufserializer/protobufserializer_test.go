package protobufserializer_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AlexSH61/homework_basic/hw09_serialize/bookpb"
	"github.com/AlexSH61/homework_basic/hw09_serialize/protobufserializer"
)

func TestProtoSliceSerDes(t *testing.T) {
	testBookSlice := []*bookpb.Book{
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
	protobuFilePath := "test_book_slice.protobuf"
	defer func() {
		if err := os.Remove(protobuFilePath); err != nil {
			t.Error("error remove file")
		}
	}()

	protobufserializer.SerializeBookSliceToProtobuf(testBookSlice, protobuFilePath)
	_, err := os.Stat(protobuFilePath)
	assert.NoError(t, err, "not open file")

	deserializeProtoBookSlice, err := protobufserializer.DeserializeProtobufSliceToBook(protobuFilePath)
	assert.NoError(t, err, "error, deserialize")
	assert.Equal(t, testBookSlice, deserializeProtoBookSlice)
}
