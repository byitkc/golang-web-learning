package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func bookPages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}