package booksd

import (
	"encoding/json"

	// "github.com/AlexSH61/homework_basic/hw09_serialize/booksd/bookproto/proto"
)

type JSONMarshaler interface {
	MarshalJSON() ([]byte, error)
}

type JSONUnmarshaler interface {
	UnmarshalJSON(data []byte) error
}

type Book struct {
	ID     int     `json:"ID"`
	Title  string  `json:"Title"`
	Author string  `json:"Author"`
	Year   int     `json:"Year"`
	Size   int     `json:"Size"`
	Rate   float64 `json:"Rate"`
}

func (b *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID     int     `json:"ID"`
		Title  string  `json:"Title"`
		Author string  `json:"Author"`
		Year   int     `json:"Year"`
		Size   int     `json:"Size"`
		Rate   float64 `json:"Rate"`
	}{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
		Size:   b.Size,
		Rate:   b.Rate,
	})
}

func (b *Book) UnmarshalJSON(data []byte) error {
	var temp struct {
		ID     int     `json:"ID"`
		Title  string  `json:"Title"`
		Author string  `json:"Author"`
		Year   int     `json:"Year"`
		Size   int     `json:"Size"`
		Rate   float64 `json:"Rate"`
	}
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}

	b.ID = temp.ID
	b.Title = temp.Title
	b.Author = temp.Author
	b.Year = temp.Year
	b.Size = temp.Size
	b.Rate = temp.Rate

	return nil
}
