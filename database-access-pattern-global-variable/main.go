package main

import (
	"fmt"
	"log"
	"net/http"

	"bookstore.ngyngcphu.com/config"
	"bookstore.ngyngcphu.com/models"
	_ "github.com/lib/pq"
)

func booksIndex(w http.ResponseWriter, r *http.Request) {
	bks, err := models.AllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}

func main() {
	err := config.InitDB("postgres://postgres:1234567890@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer config.CloseDB()

	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":3000", nil)
}
