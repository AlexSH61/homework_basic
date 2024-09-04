package repository

import (
	"context"
	"fmt"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Product interface {
	InsertProduct(ctx context.Context, product *model.Product) (int, error)
	UpdateProductbyID(ctx context.Context, id int) error
	DeleteProductbyID(ctx context.Context, id int) error
}
type ProdRepo struct {
	DB *pgxpool.Pool
}

func NewProduct(db *pgxpool.Pool) *ProdRepo {
	return &ProdRepo{
		DB: db,
	}
}

func (db *ProdRepo) InsertProduct(ctx context.Context, product *model.Product) (int, error) {
	var id int
	tx, err := db.DB.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("error in start transaction: %w", err)
	}
	defer tx.Rollback(ctx)
	query := `insert  into products(name, price)
	values($1,$2)
	returning id`
	err = tx.QueryRow(ctx, query, product.Name, product.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("transaction failed: %w", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return 0, fmt.Errorf("committing failed: %w", err)
	}
	product.ID = id
	return id, err
}

func (db *ProdRepo) UpdateProductbyID(ctx context.Context, id int, product *model.Product) error {
	tx, err := db.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("start transaction failed: %w", err)
	}
	defer tx.Rollback(ctx)
	query := `UPDATE products SET name = $1, price = $2 WHERE id = $3`
	_, err = tx.Exec(ctx, query, product.Name, product.Price, id)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("committing failed: %w", err)
	}
	return nil
}

func (db *ProdRepo) DeleteProductbyID(ctx context.Context, id int) error {
	tx, err := db.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("start transaction failed: %w", err)
	}
	defer tx.Rollback(ctx)
	query := `delete from products where id = $1`
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("transaction delete failed :%w", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("transaction committing failed: %w", err)
	}
	return nil
}
