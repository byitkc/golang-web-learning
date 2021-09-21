package main

import (
	"net/http"

	"github.com/byitkc/golang-web-learning/web/router"
)

const listenPort = ":8080"

func main() {
	r := router.BaseRouter()
	http.ListenAndServe(listenPort, r)
}
