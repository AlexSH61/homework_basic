package protobufserializer

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"

	"github.com/AlexSH61/homework_basic/hw09_serialize/bookpb"
)

func SerializeBookToProtobuf(b *bookpb.Book, filename string) {
	out, err := proto.Marshal(b)
	if err != nil {
		fmt.Println("Error during Protobuf serialization")
		return
	}

	if err := os.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("Failed to write Protobuf:", err)
	}
}

func DeserializeProtobufToBook(filename string) (*bookpb.Book, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var b bookpb.Book
	if err := proto.Unmarshal(data, &b); err != nil {
		return nil, err
	}

	return &b, nil
}
func SerializeBookSliceToProtobuf(books []*bookpb.Book, filename string) {
	bookList := &bookpb.BookList{Books: books}
	out, err := proto.Marshal(bookList)
	if err != nil {
		fmt.Println("Error during Protobuf serialization")
		return
	}

	if err := os.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("Failed to write Protobuf:", err)
	}
}

func DeserializeProtobufSliceToBook(filename string) ([]*bookpb.Book, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var bookList bookpb.BookList
	if err := proto.Unmarshal(data, &bookList); err != nil {
		return nil, err
	}

	return bookList.Books, nil
}
