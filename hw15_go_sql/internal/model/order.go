package model

import "time"

type Order struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	OrderDate   time.Time `json:"order_date"`
	TotalAmount float64   `json:"total_amount"`
}

func NewOrder(id int, userID int, totalAmount float64, ordersdate time.Time) *Order {
	return &Order{
		ID:          id,
		UserID:      userID,
		OrderDate:   ordersdate,
		TotalAmount: totalAmount,
	}
}
