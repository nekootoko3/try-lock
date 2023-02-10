package main

import (
	"fmt"
	"try-lock/util"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("creating database started")

	util.CreateUniqDB()

	fmt.Println("creating database done")

	fmt.Println("creating tables started")

	db, err := util.ConnectUniqDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec("CREATE TABLE bills (id SERIAL PRIMARY KEY)"); err != nil {
		panic(err)
	}
	if _, err := db.Exec("CREATE TABLE payments (id SERIAL PRIMARY KEY, bill_id INTEGER NOT NULL UNIQUE REFERENCES bills)"); err != nil {
		panic(err)
	}

	fmt.Println("creating tables done")

	fmt.Println("seeding started")

	for i := 0; i < 10; i++ {
		if _, err := db.Exec("INSERT INTO bills (id) VALUES ($1)", i); err != nil {
			panic(err)
		}
	}

	fmt.Println("seeding done!")
}
