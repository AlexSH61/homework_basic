package protobufserializer

import (
	"fmt"

	"github.com/AlexSH61/homework_basic/hw09_serialize/bookpb"
	"google.golang.org/protobuf/proto"
)

func SerializeBookToProtobuf(b *bookpb.Book) ([]byte, error) {
	out, err := proto.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("error during Protobuf serialization: %w", err)
	}

	return out, nil
}

func DeserializeProtobufToBook(data []byte) (*bookpb.Book, error) {
	var desBook bookpb.Book
	if err := proto.Unmarshal(data, &desBook); err != nil {
		return nil, fmt.Errorf("error during Protobuf deserialization: %w", err)
	}

	return &desBook, nil
}

func SerializeBookSliceToProtobuf(books []*bookpb.Book) ([]byte, error) {
	bookList := &bookpb.BookList{Books: books}
	out, err := proto.Marshal(bookList)
	if err != nil {
		return nil, fmt.Errorf("error during Protobuf serialization: %w", err)
	}
	return out, nil
}

func DeserializeProtobufSliceToBook(data []byte) ([]*bookpb.Book, error) {
	var bookList bookpb.BookList
	if err := proto.Unmarshal(data, &bookList); err != nil {
		return nil, fmt.Errorf("error during Protobuf deserialization: %w", err)
	}
	return bookList.Books, nil
}
