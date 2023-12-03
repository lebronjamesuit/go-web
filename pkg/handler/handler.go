package handler

import (
	"errors"
	"fmt"
	"jamesvo.uk/website/pkg/render"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var home string = "home.page.tmpl"
	render.RenderTemplate(w, home)

}

// Tach func ra rieng cho home
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	var about string = "about.page.tmpl"
	render.RenderTemplate(w, about)
}

// product handler
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	product := "product.page.tmpl"
	render.RenderTemplate(w, product)
}

// Service device | f float, d int, need to read documentation
func DevideHandler(w http.ResponseWriter, r *http.Request) {
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
