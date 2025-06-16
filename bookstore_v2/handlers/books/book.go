package books

import (
	"fmt"
	"net/http"

	"bookstore_v2.ngyngcphu.com/config"
	"bookstore_v2.ngyngcphu.com/handlers"
	"bookstore_v2.ngyngcphu.com/models"
)

func Index(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := env.GetDB()
		if err != nil {
			handlers.HTTPError(w, "getting DB", err)
			return
		}

		bks, err := models.AllBooks(db)
		if err != nil {
			handlers.HTTPError(w, "fetching books", err)
			return
		}

		for _, bk := range bks {
			fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
		}
	}
}
