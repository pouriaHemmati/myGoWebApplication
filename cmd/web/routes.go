package main

import (
	"github.com/bmizerany/pat"
	"myGoWebApplication/pkg/config"
	"myGoWebApplication/pkg/handlers"
	"net/http"
)

func routes(c *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
