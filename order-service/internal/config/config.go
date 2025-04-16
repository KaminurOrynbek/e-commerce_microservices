package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(connStr string) (*Storage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &Storage{DB: db}, nil
}
