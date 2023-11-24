package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8888"

// Tach func ra rieng cho home

func about(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This is about page")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page")
}

// Tach func ra rieng cho about

func main() {

	fmt.Println("inside main")

	// Muc tieu la request response
	/* http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "json/application")

		n, err := fmt.Fprintf(w, "abc strig ")

		fmt.Println(fmt.Sprintf("number of byte %d", n))
		fmt.Println(err)

	})
	*/
	http.HandleFunc("/about", about)

	http.HandleFunc("/home", home)

	// start web server , if error then ignore it _
	_ = http.ListenAndServe(portNumber, nil)

}
