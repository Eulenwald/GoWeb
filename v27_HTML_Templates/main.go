package main

import (
	"fmt"
	"net/http"
	"text/template"
)


const portNumber = 8080

// Handler for the startpage
func Home(w http.ResponseWriter, r *http.Request) {
	//n, err := fmt.Fprintf(w, "Hello, 世界")
	//fmt.Fprintf(w, "Hello, 世界")
	renderATemplate(w, "home-page.tpml")
}


// Handler for the about-page
func About(w http.ResponseWriter, r *http.Request) {
	//n, err := fmt.Fprintf(w, "How I'am")
	//fmt.Fprintf(w, "How I'am")
	renderATemplate(w, "about-page.tpml")
}

func renderATemplate(w http.ResponseWriter, tpml string) {
	rTemp, err := template.ParseFiles("./templates/" + tpml)

	if err != nil {
		fmt.Println("An error is happen. Exactly:", err)
	}
	err = rTemp.Execute(w, nil)

	if(err != nil) {
		fmt.Println("An error is happen. Exactly:", err)
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("GoWeb started on port %d", portNumber)

	_ = http.ListenAndServe(fmt.Sprintf(":%d", portNumber), nil)

}
