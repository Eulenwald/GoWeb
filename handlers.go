package main

import "net/http"

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