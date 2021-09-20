package main

import (
	"log"
	"net/http"

	"github.com/byitkc/goland-web-learning/pkg/config"
	"github.com/byitkc/goland-web-learning/pkg/handlers"
	"github.com/byitkc/goland-web-learning/pkg/render"
)

const listenPort = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Printf("Starting application on port %s\n", listenPort)
	_ = http.ListenAndServe(listenPort, nil)
}
