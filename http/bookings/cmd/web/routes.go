package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sojoudian/golangCheatsheet/http/booking/pkg/config"
	"github.com/sojoudian/golangCheatsheet/http/booking/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./statics/"))
	mux.Handle("/statics/*", http.StripPrefix("/statics", fileServer))

	return mux
}
