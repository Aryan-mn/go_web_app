package main

import (
	"fmt"
	"github.com/Aryan-mn/go_web_app/internal/config"
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestRoutes(t *testing.T){
	var app config.AppConfig
	
	mux := routes(&app)

	switch v:= mux.(type) {
	case *chi.Mux:
		//nothing
	default:
		t.Error(fmt.Sprintf("type is not *chi.mux, type is %t" , v))

	}
}
