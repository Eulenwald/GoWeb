package main

import (
	"net/http"

	"github.com/Eulenwald/GoWeb/pkg/config"
	"github.com/Eulenwald/GoWeb/pkg/handler"
	"github.com/Eulenwald/GoWeb/pkg/patty"
)


func routes(pAppConfig *config.AppConfig) http.Handler {
	mux := patty.New()

	pRepository := handler.GetRepository()

	mux.Get("/", http.HandlerFunc(pRepository.Home))
	mux.Get("/about", http.HandlerFunc(pRepository.About))

	return mux
}