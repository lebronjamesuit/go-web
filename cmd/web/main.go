package main

import (
	"fmt"
	"log"
	"net/http"

	"jamesvo.uk/website/pkg/config"
	"jamesvo.uk/website/pkg/handler"
	"jamesvo.uk/website/pkg/render"
)

const portNumber = ":8888"

func main() {

	appConfig := config.AppConfig{}
	allCaches, err := render.CreateTemplateCaches()
	if err != nil {
		log.Fatal("CreateAllTemplate has issues ")
	}
	appConfig.MyCaches = allCaches

	// Give render the access to the same Instace of AppConfig as main has
	render.NewTemplate(&appConfig)

	http.HandleFunc("/about", handler.AboutHandler)

	http.HandleFunc("/home", handler.HomeHandler)

	http.HandleFunc("/devide", handler.DevideHandler)

	http.HandleFunc("/product", handler.ProductHandler)

	// start web server , if error then ignore it _

	fmt.Print("start web server ok at localhost:8888")
	_ = http.ListenAndServe(portNumber, nil)

}
