package main

import (
	"fmt"
	"net/http"

	"github.com/Eulenwald/GoWeb/pkg/handler"
)

const portNumber = 8080

func main() {

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Printf("GoWeb started on port %d", portNumber)

	_ = http.ListenAndServe(fmt.Sprintf(":%d", portNumber), nil)

}
