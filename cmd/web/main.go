package main

import (
	"fmt"
	"net/http"

	"jamesvo.uk/website/pkg/handler"
)

const portNumber = ":8888"

func main() {

	http.HandleFunc("/about", handler.AboutHandler)

	http.HandleFunc("/home", handler.HomeHandler)

	http.HandleFunc("/devide", handler.DevideHandler)

	http.HandleFunc("/product", handler.ProductHandler)

	// start web server , if error then ignore it _

	fmt.Print("start web server ok at localhost:8888")
	_ = http.ListenAndServe(portNumber, nil)

}
