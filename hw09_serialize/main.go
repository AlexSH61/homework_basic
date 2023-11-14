package main

import (
	"encoding/json"
	"fmt"

	booksd "github.com/AlexSH61/homework_basic/hw09_serialize/booksd"
	"github.com/AlexSH61/homework_basic/hw09_serialize/booksd/bookproto"
	"github.com/AlexSH61/homework_basic/hw09_serialize/booksd/bookproto/proto"
)

func main() {
	b1 := &booksd.Book{ID: 1, Title: "Book1", Author: "Author1", Year: 2022, Size: 200, Rate: 4.5}
	b2 := &booksd.Book{ID: 2, Title: "Book2", Author: "Author2", Year: 2023, Size: 250, Rate: 4.0}
	fmt.Println(b2)
	jsonBytes, err := json.Marshal(b1)
	if err != nil {
		fmt.Printf("JSON marshal error: %s\n", err.Error())
		return
	}

	var deserializedBooks []*booksd.Book
	err = json.Unmarshal(jsonBytes, &deserializedBooks)
	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return
	}
	fmt.Println("\nDeserialized from JSON:")
	for _, book := range deserializedBooks {
		fmt.Printf("%+v\n", *book)
	}
	protoBookSlice := []*proto.Book{
		{ID: 1, Title: "Book 1", Author: "Author 1", Year: 2020, Size: 200, Rate: 4.0},
		{ID: 2, Title: "Book 2", Author: "Author 2", Year: 2021, Size: 250, Rate: 4.5},
		{ID: 3, Title: "Book 3", Author: "Author 3", Year: 2022, Size: 300, Rate: 5.0},
	}

	serializedData, err := bookproto.SerializeBooks(protoBookSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(serializedData)

	deserializedBooks, err := bookproto.DeserializeBooks(serializedData)
	if err != nil {
		fmt.Println(err)
	}
	for _, book := range deserializedBooks {
		fmt.Printf("%+v\n", *book)
		// }
	}
}
