package main

import (
	"log"
	"net/http"

	"bookstore_v2.ngyngcphu.com/config"
	"bookstore_v2.ngyngcphu.com/handlers/books"
	_ "github.com/lib/pq"
)

func main() {
	env := config.NewEnv()
	err := env.InitDB("postgres://postgres:1234567890@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer env.CloseDB()

	http.HandleFunc("/books", books.Index(env))
	http.ListenAndServe(":3000", nil)
}
