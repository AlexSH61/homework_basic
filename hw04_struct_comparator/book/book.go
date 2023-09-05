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

// func (b Book) SetID(id int) {
// 	b.id = id
// }

func (b Book) GetID() int {
	return b.id
}

// func (b *Book) SetTitle(title string) {
// 	b.title = title
// }

func (b Book) GetTitle() string {
	return b.title
}

// func (b *Book) SetAuthor(author string) {
// 	b.author = author
// }

func (b Book) GetAuthor() string {
	return b.author
}

//	func (b *Book) SetYear(year int) {
//		b.year = year
//	}
func (b Book) GetYear() int {
	return b.year
}
func (b Book) GetSize() int {
	return b.size
}

// func (b *Book) SetSize(year int) {
// 	b.year = year
// }

// func (b *Book) SetRate(rating float32) {
// 	b.rating = rating
// }

func (b Book) GetRating() float32 {
	return b.rating
}
