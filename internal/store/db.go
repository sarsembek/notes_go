package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type Store struct {
	db *sql.DB
}

func NewStore(dataSourceName string) *Store {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return &Store{db: db}
}

// Здесь будут функции для работы с базой данных
