package main

import (
	"net/http"
)

const portNumber = ":8888"

func main() {

	http.HandleFunc("/about", aboutHandler)

	http.HandleFunc("/home", homeHandler)

	http.HandleFunc("/devide", devideHandler)

	http.HandleFunc("/product", productHandler)

	// start web server , if error then ignore it _
	_ = http.ListenAndServe(portNumber, nil)

}
