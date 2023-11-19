package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Eulenwald/GoWeb/pkg/config"
	"github.com/Eulenwald/GoWeb/pkg/handler"
	"github.com/Eulenwald/GoWeb/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	// No cache no tool => game over
	tc, err  := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("can not create the cache!")
	}

	// save the cache in app and decide whether the cache should be used
	app.TemplateCache = tc
	app.UseCache = false

	// linked the app module to "handler" and "render"
	pRepository := handler.NewRepository(&app)	
	handler.NewHandler(pRepository)
	render.NewTemplate(&app)

	fmt.Printf("GoWeb started on port %v", portNumber)

	//_ = http.ListenAndServe(fmt.Sprintf(":%d", portNumber), nil)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app)	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Server konnte nicht starten")
	}

}
