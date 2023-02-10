package main

import (
	"fmt"
	"log"
	"try-lock/util"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("creating database started")

	util.CreatePessimisticDB()

	fmt.Println("creating database done")

	fmt.Println("seeding started")

	db2, err := util.ConnectPessimisticDB()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := db2.Exec("CREATE TABLE bills (id SERIAL PRIMARY KEY, paid_at timestamp)"); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		if _, err := db2.Exec("INSERT INTO bills (id) VALUES ($1)", i); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("seeding done!")
}
