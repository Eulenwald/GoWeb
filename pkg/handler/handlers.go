package handler

import (
	"net/http"

	"github.com/Eulenwald/GoWeb/pkg/render"
)

// Handler for the startpage
func Home(w http.ResponseWriter, r *http.Request) {
	//n, err := fmt.Fprintf(w, "Hello, 世界")
	//fmt.Fprintf(w, "Hello, 世界")
	render.RenderATemplate(w, "home-page.tpml")
}

// Handler for the about-page
func About(w http.ResponseWriter, r *http.Request) {
	//n, err := fmt.Fprintf(w, "How I'am")
	//fmt.Fprintf(w, "How I'am")
	render.RenderATemplate(w, "about-page.tpml")
}
