package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hakem26/hello_web_go/pkg/config"
	"github.com/hakem26/hello_web_go/pkg/handlers"
	"github.com/hakem26/hello_web_go/pkg/render"
)

const portNumber = ":8080"

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

	fmt.Printf("%s", fmt.Sprintf("The Server Started On Port %s \n", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}