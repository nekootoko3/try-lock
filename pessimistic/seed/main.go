package main

import (
	"fmt"
	"try-lock/util"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("creating database started")

	db, err := util.ConnectPostgresDB()
	if err != nil {
		panic(err)
	}
	if _, err := db.Exec("DROP DATABASE IF EXISTS pessimistic"); err != nil {
		panic(err)
	}
	if _, err := db.Exec("CREATE DATABASE pessimistic"); err != nil {
		panic(err)
	}
	if err := db.Close(); err != nil {
		panic(err)
	}

	fmt.Println("creating database done")

	fmt.Println("seeding started")

	db2, err := util.ConnectPessimisticDB()
	if err != nil {
		panic(err)
	}
	if _, err := db2.Exec("CREATE TABLE bills (id SERIAL PRIMARY KEY, paid_at timestamp)"); err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		if _, err := db2.Exec("INSERT INTO bills (id) VALUES ($1)", i); err != nil {
			panic(err)
		}
	}

	fmt.Println("seeding done!")
}
