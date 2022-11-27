package main

import (
	"net/http"

	"github.com/fontexd/go/health/pkg/config"
	"github.com/fontexd/go/health/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.Appconfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(noSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/environmets", handlers.Repo.Environments)

	fileServer := http.FileServer(http.Dir("./Static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
