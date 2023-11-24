package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8888"

// No template
func productHandler(w http.ResponseWriter, r *http.Request) {

	pTemplate, err := template.ParseFiles("./template/product.page.tmpl")
	if err == nil {
		fmt.Println("Not found page product")
	}

	pTemplate.Execute(w, nil)
}

func main() {

	http.HandleFunc("/about", aboutHandler)

	http.HandleFunc("/home", homeHandler)

	http.HandleFunc("/devide", devideHandler)

	http.HandleFunc("/product", productHandler)

	// start web server , if error then ignore it _
	_ = http.ListenAndServe(portNumber, nil)

}
