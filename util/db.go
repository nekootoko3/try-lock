package util

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectPostgresDB() (*sql.DB, error) {
	return sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
}

func ConnectPessimisticDB() (*sql.DB, error) {
	return sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=pessimistic sslmode=disable")
}
