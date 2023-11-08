package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func renderATemplate(w http.ResponseWriter, tpml string) {
	rTemp, err := template.ParseFiles("./templates/" + tpml)

	if err != nil {
		fmt.Println("An error is happen. Exactly:", err)
	}
	err = rTemp.Execute(w, nil)

	if err != nil {
		fmt.Println("An error is happen. Exactly:", err)
	}
}
