package jsonserializer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AlexSH61/homework_basic/hw09_serialize/book"
)

func SerializeBookToJSON(b []book.Book, filename string) {
	serializeBook, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error during JSON serialization")
	}

	if err = os.WriteFile(filename, serializeBook, 0644); err != nil {
		fmt.Println("error writing JSON file")
	}
}

func DeserializeJSONToBook(filename string) ([]book.Book, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []book.Book{}, err
	}

	var books []book.Book
	if err = json.Unmarshal(data, &books); err != nil {
		return []book.Book{}, err
	}

	return books, nil
}
