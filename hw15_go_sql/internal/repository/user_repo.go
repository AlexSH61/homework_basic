package repository

import (
	"context"
	"fmt"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type User interface {
	InsertUser(ctx context.Context, user *model.User) (int, error)
	UpdateUserbyID(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, ID int) (*model.User, error)
	DeleteUsersById(ctx context.Context, ID int) error
}

type DBUserRepo struct {
	DB *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) *DBUserRepo {
	return &DBUserRepo{
		DB: db,
	}
}

func (dbu *DBUserRepo) InsertUser(ctx context.Context, user *model.User) (int, error) {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var ID int
	err := dbu.DB.QueryRow(ctx, query, user.Name, user.Email, user.PassWord).Scan(&ID)
	if err != nil {
		return 0, fmt.Errorf("query execution error: %w", err)
	}
	return ID, nil
}

func (dbu *DBUserRepo) GetUserByID(ctx context.Context, ID int) (*model.User, error) {
	user := &model.User{}
	err := dbu.DB.QueryRow(ctx, `
	SELECT name, email, password FROM users WHERE id = $1 LIMIT 1;`, ID).Scan(&user.Name, &user.Email, &user.PassWord)
	if err != nil {
		return nil, fmt.Errorf("query execution error: %w ", err)
	}

	user.ID = ID
	return user, nil
}

func (dbu *DBUserRepo) UpdateUserByID(ctx context.Context, user *model.User) error {
	query := `UPDATE users
	          SET name = $1, email = $2, password = $3
	          WHERE id = $4`

	_, err := dbu.DB.Exec(ctx, query, user.Name, user.Email, user.PassWord, user.ID)
	if err != nil {
		return fmt.Errorf("query execution error: %w", err)
	}
	return nil
}

func (dbu *DBUserRepo) DeleteUsersById(ctx context.Context, ID int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := dbu.DB.Exec(ctx, query, ID)
	if err != nil {
		return fmt.Errorf("transaction process failed: %w", err)
	}
	return nil
}
