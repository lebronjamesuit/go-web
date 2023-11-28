package main

import (
	"fmt"
	"net/http"

	"jamesvo.uk/website/pkg/handler"
)

const portNumber = ":8888"

func main() {

	testcho()

	http.HandleFunc("/about", handler.AboutHandler)

	http.HandleFunc("/home", handler.HomeHandler)

	http.HandleFunc("/devide", handler.DevideHandler)

	http.HandleFunc("/product", handler.ProductHandler)

	// start web server , if error then ignore it _

	fmt.Print("start web server ok at localhost:8888")
	_ = http.ListenAndServe(portNumber, nil)

}

func testcho() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

}
