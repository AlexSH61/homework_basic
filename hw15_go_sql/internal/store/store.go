package store

import (
	"database/sql"

	"go.uber.org/zap"
)

type Store struct {
	config *Config
	db     *sql.DB
	log    *zap.SugaredLogger
}

func New(config *Config, log *zap.SugaredLogger) *Store {
	return &Store{
		config: config,
		log:    log}
}
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DataBaseURI)
	if err != nil {
		s.log.Errorf("COnnection failed %v\n")
		return err
	}
	if err := db.Ping(); err != nil {
		if s.log != nil {
			s.log.Errorf("Ping failed %v\n", err)
		}
	}
	s.db = db
	return nil
}
func (s *Store) Close() {
	s.db.Close()
}
