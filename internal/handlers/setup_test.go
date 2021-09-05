package handlers

import (
	"encoding/gob"
	"fmt"
	"github.com/Aryan-mn/go_web_app/internal/config"
	"github.com/Aryan-mn/go_web_app/internal/model"
	"github.com/Aryan-mn/go_web_app/internal/render"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
	"log"
	"net/http"
	"path/filepath"
	"time"
	"html/template"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"
var functions = template.FuncMap{}


func getRoutes() http.Handler{
	//stuff that i put in the session
	gob.Register(model.Reservation{})

	// change to true when in production
	app.InProduction = true

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := CreatTestTemplateCache()
	if err !=nil{
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := NewRepo(&app)
	NewHandlers(repo)

	render.NewTemplate(&app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(sessionLoad)

	mux.Get("/", Repo.Index)
	mux.Get("/about", Repo.About)
	mux.Get("/generals-quarters", Repo.Generals)
	mux.Get("/majors-suite", Repo.Majors)


	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Get("/search-availability-json", Repo.AvailabilityJson)
	mux.Get("/reservation-summary", Repo.ReservationSummary)


	mux.Get("/contact", Repo.Contact)


	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)


	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux


	//http.HandleFunc("/", Repo.Index)
	//http.HandleFunc("/about", Repo.About)

}


func NoSurf(next http.Handler)http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}


// sessionLoad loads and saves the session on every request
func sessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}

func CreatTestTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages,err := filepath.Glob(fmt.Sprintf("%s/*.gohtml", pathToTemplates))
	if err !=nil{
		return myCache, err
	}

	for _,page := range pages {
		name :=filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err !=nil{
			return myCache,err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.gohtml", pathToTemplates))
		if err !=nil{
			return myCache,err
		}
		if len(matches) >0 {
			ts,err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.gohtml", pathToTemplates))
			if err !=nil{
				return myCache,err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}