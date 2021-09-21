package router

import (
	"github.com/byitkc/golang-web-learning/web/render"
	"github.com/gorilla/mux"
)

func BaseRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", render.HomePage)
	r.HandleFunc("/books/{title}/page/{page}", render.BookPages)

	return r
}
