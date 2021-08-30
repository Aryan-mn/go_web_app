package main

import (
	"github.com/Aryan-mn/go_web_app/pkg/config"
	"github.com/Aryan-mn/go_web_app/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler{
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(noSurf)

	mux.Get("/", handlers.Repo.Index)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
