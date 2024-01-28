package db

import (
	"fmt"

	"github.com/jackc/pgx"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/config"
)

type DataBase struct {
	conn *pgx.Conn
}

func NewDataBase(cfg *config.Setting) (*DataBase, error) {

	connConfig, err := pgx.ParseConnectionString(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgNameBD))
	if err != nil {
		return nil, fmt.Errorf("error parsing connection config: %v", err)
	}

	conn, err := pgx.Connect(connConfig)
	if err != nil {
		return nil, fmt.Errorf("error parsing connection config: %v", err)
	}
	return &DataBase{conn: conn}, nil
}

func (db *DataBase) Close() {
	db.conn.Close()
}

func (db *DataBase) BeginTransaction() (*pgx.Tx, error) {
	return db.conn.Begin()
}
