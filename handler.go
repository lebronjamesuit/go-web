package main

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var home string = "about.page.tmpl"
	renderTemplate(w, home)

}

// Tach func ra rieng cho home
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	var about string = "about.page.tmpl"
	renderTemplate(w, about)
}

// product handler
func productHandler(w http.ResponseWriter, r *http.Request) {
	product := "product.page.tmpl"
	renderTemplate(w, product)
}

func renderTemplate(w http.ResponseWriter, fileName string) {

	pointerTem, _ := template.ParseFiles("./template/" + fileName)

	err := pointerTem.Execute(w, nil)
	if err != nil {
		fmt.Println("Error generating template")
		return
	}

}

// Service device | f float, d int, need to read documentation
func devideHandler(w http.ResponseWriter, r *http.Request) {

	z, err := devideTwoNumbers(10, 0)

	fmt.Fprintf(w, fmt.Sprintf("Result is %f, and error is %d", z, err))

}

func devideTwoNumbers(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot devide if the y is 0")
		return 0, err
	}

	z := x / y
	return z, nil

}
