package bookproto

import (
	// "github.com/AlexSH61/homework_basic/hw09_serialize/booksd"
	"github.com/AlexSH61/homework_basic/hw09_serialize/booksd/bookproto/proto"
)

func SerializeBooks(books []*proto.Book) ([]byte, error) {
	booksList := &proto.Bookslist{Books: books}
	return proto.Marshal(booksList)
}

func DeserializeBooks(data []byte) ([]*proto.Book, error) {
	var booksList Bookslist
	err := proto.Unmarshal(data, &booksList)
	if err != nil {
		return nil, err
	}
	return booksList.Books, nil
}
