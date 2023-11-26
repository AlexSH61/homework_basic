package main

import (
	"encoding/json"
	"fmt"

	"github.com/AlexSH61/homework_basic/hw09_serialize/book"
	"github.com/AlexSH61/homework_basic/hw09_serialize/bookpb"
	"github.com/AlexSH61/homework_basic/hw09_serialize/jsonserializer"
	"github.com/AlexSH61/homework_basic/hw09_serialize/protobufserializer"
)

func main() {
	exampleBook := &book.Book{
		ID:     1,
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan",
		Year:   2015,
		Size:   356,
		Rate:   4.5,
	}
	sampleBooks := &[]book.Book{
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

	protoBufBook := &bookpb.Book{
		ID:     2,
		Title:  "Clean Code: A Handbook of Agile Software Craftsmanship",
		Author: "Robert C. Martin",
		Year:   2008,
		Size:   464,
		Rate:   4.8,
	}
	serToSliceByte, err := jsonserializer.SerializeBookToJSON(*sampleBooks)
	if err != nil {
		fmt.Println("Error during JSON serialization:", err)
		return
	}
	fmt.Println(string(serToSliceByte))

	booksFromJSON, err := jsonserializer.DeserializeJSONToBook(serToSliceByte)
	if err != nil {
		fmt.Println("Error during JSON deserialization:", err)
		return
	}
	fmt.Println(booksFromJSON)

	serSliceBooks, err := protobufserializer.SerializeBookToProtobuf(protoBufBook)
	if err != nil {
		fmt.Println("Error during serialization to protobuf")
		return
	}
	fmt.Println(serSliceBooks)

	desSliceBooks, err := protobufserializer.DeserializeProtobufToBook(serSliceBooks)
	if err != nil {
		fmt.Println("Error during deserialization from protobuf")
		return
	}
	fmt.Println(desSliceBooks)

	serProtoBook, err := protobufserializer.SerializeBookToProtobuf(protoBufBook)
	if err != nil {
		fmt.Println("Error during serialization using protobuf")
		return
	}
	fmt.Println(serProtoBook)

	desProtoBook, err := protobufserializer.DeserializeProtobufToBook(serProtoBook)
	if err != nil {
		fmt.Println("Error during deserialization using protobuf")
		return
	}
	fmt.Println(desProtoBook)

	serializedData, err := json.Marshal(exampleBook)
	if err != nil {
		fmt.Println("Error during JSON serialization:", err)
		return
	}
	fmt.Println("Serialized data:", string(serializedData))

	deserializedBook := &book.Book{}
	err = json.Unmarshal(serializedData, deserializedBook)
	if err != nil {
		fmt.Println("Error during JSON deserialization:", err)
		return
	}
	fmt.Printf("Deserialized book: %+v\n", deserializedBook)
}
