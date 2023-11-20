package protobufserializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"

	"github.com/AlexSH61/homework_basic/hw09_serialize/bookpb"
)

func SerializeBookToProtobuf(b *bookpb.Book, filename string) error {
	out, err := proto.Marshal(b)
	if err != nil {
		return fmt.Errorf("error during Protobuf serialization: %w", err)
	}

	return os.WriteFile(filename, out, 0600)
}

func DeserializeProtobufToBook(filename string) (*bookpb.Book, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var b bookpb.Book
	if err = proto.Unmarshal(data, &b); err != nil {
		return nil, fmt.Errorf("error during Protobuf deserialization: %w", err)
	}

	return &b, nil
}

func SerializeBookSliceToProtobuf(books []*bookpb.Book, filename string) error {
	bookList := &bookpb.BookList{Books: books}
	out, err := proto.Marshal(bookList)
	if err != nil {
		return fmt.Errorf("error during Protobuf serialization: %w", err)
	}

	return os.WriteFile(filename, out, 0600)
}

func DeserializeProtobufSliceToBook(filename string) ([]*bookpb.Book, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var bookList bookpb.BookList
	if err = proto.Unmarshal(data, &bookList); err != nil {
		return nil, fmt.Errorf("error during Protobuf deserialization: %w", err)
	}

	return bookList.Books, nil
}
