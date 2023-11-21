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

	serToSliceByte, err := jsonserializer.SerializeBookToJSON(*sampleBook)
	if err != nil {
		fmt.Println("Error during JSON serialization:", err)
		return
	}

	fmt.Println(string(serToSliceByte))

	BooksFromJSON, err := jsonserializer.DeserializeJSONToBook(serToSliceByte)
	if err != nil {
		fmt.Println("Error during JSON deserialization:", err)
		return
	}

	fmt.Println(BooksFromJSON)

	serSliceBooks, err := protobufserializer.SerializeBookToProtobuf(protobufBook)
	if err != nil {
		fmt.Println("err при сериализации слайса в протофайле")
		return
	}
	fmt.Println(serSliceBooks)

	desSliceBooks, err := protobufserializer.DeserializeProtobufToBook(serToSliceByte)
	if err != nil {
		fmt.Println("err при десериализации слайса в протофайле")
		return
	}
	fmt.Println(desSliceBooks)

	serProtoBook, err := protobufserializer.SerializeBookToProtobuf(protobufBook)
	if err != nil {
		fmt.Println("err при сериалзиции с использованием протофайлом")
		return
	}
	fmt.Println(serProtoBook)

	desProtoBook, err := protobufserializer.DeserializeProtobufToBook(serProtoBook)
	if err != nil {
		fmt.Println("err при десериалзиции с использованием протофайлом")
		return
	}
	fmt.Println(desProtoBook)
}
