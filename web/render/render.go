package render

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HomePage is the renderer for the Home page
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the future home of farts!")
}

// Books is the renderer for the books pages
func BookPages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}
