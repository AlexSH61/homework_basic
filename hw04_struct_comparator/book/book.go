package book

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rating float32
}

// create constructor book struct.
func NewBook(id int, title, author string, year, size int, rating float32) *Book {
	return &Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rating: rating,
	}
}

func (b *Book) ID(id int) {
	b.id = id
}
func (b *Book) GetID() int {
	return b.id
}
func (b *Book) Title(title string) {
	b.title = title
}
func (b *Book) GetTitle() string {
	return b.title
}
func (b *Book) Author(author string) {
	b.author = author
}
func (b *Book) GetAuthor() string {
	return b.author
}
func (b *Book) Year(year int) {
	b.year = year
}
func (b *Book) GetYear() int {
	return b.year
}
func (b *Book) GetSize() int {
	return b.size
}
func (b *Book) Size(year int) {
	b.year = year
}
func (b *Book) Rating(rating float32) {
	b.rating = rating
}
func (b *Book) GetRating() float32 {
	return b.rating
}
