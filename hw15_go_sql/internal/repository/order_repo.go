package repository

import (
	"context"
	"fmt"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Order interface {
	InsertOrder(ctx context.Context, order *model.Order) (int, error)
	GetOrdersByUserID(ctx context.Context, userID int) ([]model.Order, error)
	UpdateOrderByID(ctx context.Context, order *model.Order) error
	DeleteOrderByID(ctx context.Context, orderID int) error
}

type DBOrderRepo struct {
	DB *pgxpool.Pool
	// logger *zap.SugaredLogger
}

func NewOrder(db *pgxpool.Pool) *DBOrderRepo {
	return &DBOrderRepo{
		DB: db,
	}
}

func (odb *DBOrderRepo) InsertOrder(ctx context.Context, order *model.Order) (int, error) {
	var orderID int

	query := `INSERT INTO orders (user_id, orderdate, totalamount) VALUES ($1, $2, $3) RETURNING id`
	err := odb.DB.QueryRow(ctx, query, order.UserID, order.OrderDate, order.TotalAmount).Scan(&orderID)
	if err != nil {
		return 0, fmt.Errorf("query failed: %w", err)
	}
	return orderID, nil
}

func (odb *DBOrderRepo) GetOrdersByUserID(ctx context.Context, userID int) ([]model.Order, error) {
	var orders []model.Order
	query := `SELECT id, user_id, orderdate, totalamount FROM orders WHERE user_id = $1`
	rows, err := odb.DB.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error in query GetOrders: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (odb *DBOrderRepo) UpdateOrderByID(ctx context.Context, order *model.Order) error {
	query := `UPDATE orders SET user_id = $1, orderdate = $2, totalamount = $3 WHERE id = $4`
	_, err := odb.DB.Exec(ctx, query, order.UserID, order.OrderDate, order.TotalAmount, order.ID)
	if err != nil {
		return fmt.Errorf("failed to update order: %w", err)
	}
	return nil
}

func (odb *DBOrderRepo) DeleteOrderByID(ctx context.Context, orderID int) error {
	query := `DELETE FROM orders WHERE id = $1`
	_, err := odb.DB.Exec(ctx, query, orderID)
	if err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}
	return nil
}
