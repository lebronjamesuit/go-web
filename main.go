package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("hello")

	// Muc tieu la request response
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "abc strig ")

		fmt.Println(n)
		fmt.Println(err)

	})

	http.ListenAndServe(":8888", nil)

}
