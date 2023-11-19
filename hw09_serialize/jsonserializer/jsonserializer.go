package jsonserializer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/AlexSH61/homework_basic/hw09_serialize/book"
)

func SerializeBookToJson(b []book.Book, filename string) {
	serializeBook, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error during JSON serialization")
		return
	}

	if err := os.WriteFile(filename, serializeBook, 0644); err != nil {
		log.Fatalln("Failed to write JSON:", err)
	}
}
func DeserializeJsonToBook(filename string) ([]book.Book, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []book.Book{}, err
	}

	var books []book.Book
	if err := json.Unmarshal(data, &books); err != nil {
		return []book.Book{}, err
	}

	return books, nil
}
