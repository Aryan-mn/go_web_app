package main

import (
	"fmt"
	"github.com/aryan_mn/test/pkg/config"
	"github.com/aryan_mn/test/pkg/render"
	"log"
	"net/http"
	"github.com/aryan_mn/test/pkg/handlers"
)


func main() {
	var app config.AppConfig
	tc, err := render.CreatTemplateCache()
	if err !=nil{
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Index)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port :8080"))
	http.ListenAndServe(":8080", nil)
}
