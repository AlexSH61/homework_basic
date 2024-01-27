package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPassword string
	PgNameBD   string
}

type DB struct {
	pool *pgxpool.Pool
}

var CFG Config

func Init() {
	file, err := os.Open("setting.cfg")
	if err != nil {
		log.Println("error by open file setting", err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
	}
	readByte := make([]byte, stat.Size())
	file.Read(readByte)
	_, err = file.Read(readByte)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(readByte, &CFG)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func NewDB(config Config) (*DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.PgHost, config.PgPort, config.PgUser, config.PgPassword, config.PgNameBD)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return &DB{pool: pool}, nil
}

func (db *DB) Close() {
	db.pool.Close()
}

func (db *DB) ExecuteInTransaction(fn func(tx pgx.Tx) error) error {
	tx, err := db.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(context.Background())
			panic(p)
		} else if err != nil {
			_ = tx.Rollback(context.Background())
		} else {
			err = tx.Commit(context.Background())
		}
	}()

	return fn(tx)
}

func (db *DB) InsertUser(name, email, password string) error {
	return db.ExecuteInTransaction(func(tx pgx.Tx) error {
		_, err := tx.Exec(context.Background(),
			"INSERT INTO schema_store.users (name, email, password) VALUES ($1, $2, $3)", name, email, password)
		return err
	})
}

func (db *DB) UpdateUserName(userID int, newName string) error {
	return db.ExecuteInTransaction(func(tx pgx.Tx) error {
		_, err := tx.Exec(context.Background(), "UPDATE schema_store.users SET name = $1 WHERE id = $2", newName, userID)
		return err
	})
}

func (db *DB) DeleteUser(userID int) error {
	return db.ExecuteInTransaction(func(tx pgx.Tx) error {
		_, err := tx.Exec(context.Background(), "DELETE FROM schema_store.users WHERE id = $1", userID)
		return err
	})
}
