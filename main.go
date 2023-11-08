package main

import (
	"errors"
	"fmt"
	"net/http"
)


const portNumber = 8080

// Handler for the startpage
func Home(w http.ResponseWriter, r *http.Request) {
	//n, err := fmt.Fprintf(w, "Hello, 世界")
	fmt.Fprintf(w, "Hello, 世界")
}


// Handler for the about-page
func About(w http.ResponseWriter, r *http.Request) {
	//n, err := fmt.Fprintf(w, "How I'am")
	fmt.Fprintf(w, "How I'am")
}

// Handler for the about-page
func Devide(w http.ResponseWriter, r *http.Request) {
	//n, err := fmt.Fprintf(w, "How I'am")
	val, err := devide(100.00, 10.00)
	if (err != nil) {
		_, _ = fmt.Fprintf(w, "Error: an interneal error is happen! Returned: %s\n", err)
		return 
	}
	fmt.Fprintf(w, "The dived of 100 and 10 is %f\n", val)

	val, err = devide(100.00, 0.00)
	if (err != nil) {
		_, _ = fmt.Fprintf(w, "Error: an interneal error is happen! Returned: %s\n", err)
		return 
	}
	fmt.Fprintf(w, "The dived of 100 and 10 is %f\n", val)
}

func devide(x, y float64) (float64, error) {
	
	rVal := 0.0

	if (y == 0) {
		err := errors.New("Devision by Zero")
		return rVal, err
	}

	rVal = x / y
	return rVal, nil
}
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/devide", Devide)

	fmt.Printf("GoWeb started on port %d", portNumber)

	_ = http.ListenAndServe(fmt.Sprintf(":%d", portNumber), nil)

}
