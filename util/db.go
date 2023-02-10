package util

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreatePessimisticDB() {
	db, err := connectPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec("DROP DATABASE IF EXISTS pessimistic"); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec("CREATE DATABASE pessimistic"); err != nil {
		log.Fatal(err)
	}
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

func CreateUniqDB() {
	db, err := connectPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec("DROP DATABASE IF EXISTS uniq"); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec("CREATE DATABASE uniq"); err != nil {
		log.Fatal(err)
	}
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

func connectPostgresDB() (*sql.DB, error) {
	return sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
}

func ConnectPessimisticDB() (*sql.DB, error) {
	return sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=pessimistic sslmode=disable")
}

func ConnectUniqDB() (*sql.DB, error) {
	return sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=uniq sslmode=disable")
}
