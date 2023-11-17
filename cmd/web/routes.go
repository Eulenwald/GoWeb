package main

import (
	"net/http"

	"github.com/Eulenwald/GoWeb/pkg/config"
	"github.com/Eulenwald/GoWeb/pkg/handler"
	"github.com/bmizerany/pat"
)

func routes(pAppConfig *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handler.PRepository.Home))
	mux.Get("/about", http.HandlerFunc(handler.PRepository.About))

	return mux
}