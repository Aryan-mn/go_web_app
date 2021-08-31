package main

import (
	"fmt"
	"github.com/Aryan-mn/go_web_app/internal/config"
	"github.com/Aryan-mn/go_web_app/internal/handlers"
	"github.com/Aryan-mn/go_web_app/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change to true when in production
	app.InProduction = false

	session = scs.New ()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
