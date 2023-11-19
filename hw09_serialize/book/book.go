package book

type Book struct {
	ID     int32   `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int32   `json:"year"`
	Size   int32   `json:"size"`
	Rate   float32 `json:"rate"`
}
