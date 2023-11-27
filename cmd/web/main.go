package main

import (
	"net/http"

	"jamesvo.uk/website/pkg/handler"
)

const portNumber = ":8888"

func main() {

	handler.Nothing()

	http.HandleFunc("/about", handler.AboutHandler)

	http.HandleFunc("/home", handler.HomeHandler)

	http.HandleFunc("/devide", handler.DevideHandler)

	http.HandleFunc("/product", handler.ProductHandler)

	// start web server , if error then ignore it _
	_ = http.ListenAndServe(portNumber, nil)

}
