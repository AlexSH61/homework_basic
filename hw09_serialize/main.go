package main

import (
	"fmt"

	"github.com/AlexSH61/homework_basic/hw09_serialize/book"
	"github.com/AlexSH61/homework_basic/hw09_serialize/bookpb"
	"github.com/AlexSH61/homework_basic/hw09_serialize/jsonserializer"
	"github.com/AlexSH61/homework_basic/hw09_serialize/protobufserializer"
)

func main() {
	sampleBook := &[]book.Book{
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

	protobufBook := &bookpb.Book{
		ID:     2,
		Title:  "Clean Code: A Handbook of Agile Software Craftsmanship",
		Author: "Robert C. Martin",
		Year:   2008,
		Size:   464,
		Rate:   4.8,
	}
	jsonserializer.SerializeBookToJSON(*sampleBook, "books.json")
	BooksFromJSON, _ := (jsonserializer.DeserializeJSONToBook("books.json"))
	fmt.Println(BooksFromJSON)
	protobufserializer.SerializeBookToProtobuf(protobufBook, "books.protobuf")
	protobufserializer.DeserializeProtobufToBook("books.protobuf")
}
