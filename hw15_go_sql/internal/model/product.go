package model

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewProduct(id int, price float64, name string) *Product {
	return &Product{
		ID:    id,
		Price: price,
		Name:  name,
	}
}
