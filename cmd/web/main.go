package main

import (
	"fmt"
	"github.com/Aryan-mn/go_web_app/pkg/config"
	"github.com/Aryan-mn/go_web_app/pkg/handlers"
	"ggithub.com/Aryan-mn/go_web_app/pkg/render"
	"log"
	"net/http"
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

	//http.HandleFunc("/", handlers.Repo.Index)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port :8080"))
	//http.ListenAndServe(":8080", nil)
	srv := &http.Server{
		Addr: ":8080",
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatalln(err)
}
