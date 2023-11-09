package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Eulenwald/GoWeb/pkg/config"
	"github.com/Eulenwald/GoWeb/pkg/handler"
	"github.com/Eulenwald/GoWeb/pkg/render"
)

const portNumber = 8080

func main() {

	var app config.AppConfig

	tc, err  := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("can not create the cache!")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandler(repo)
	render.NewTemplate(&app)

	http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/about", handler.Repo.About)

	fmt.Printf("GoWeb started on port %d", portNumber)

	_ = http.ListenAndServe(fmt.Sprintf(":%d", portNumber), nil)

}
