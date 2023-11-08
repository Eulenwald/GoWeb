package main

import (
	"fmt"
	"net/http"
)


const portNumber = 8080



func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("GoWeb started on port %d", portNumber)

	_ = http.ListenAndServe(fmt.Sprintf(":%d", portNumber), nil)

}
