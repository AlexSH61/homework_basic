package jsonserializer

import (
	"encoding/json"
	"fmt"

	"github.com/AlexSH61/homework_basic/hw09_serialize/book"
)

func SerializeBookToJSON(b []book.Book) ([]byte, error) {
	serializeBook, err := json.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("err при сериализации")
	}
	return serializeBook, nil
}

func DeserializeJSONToBook(data []byte) ([]book.Book, error) {
	var books []book.Book
	err := json.Unmarshal(data, &books)
	if err != nil {
		return nil, fmt.Errorf("err при десериализации")
	}
	return books, nil
}
