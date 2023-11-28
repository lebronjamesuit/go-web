package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplateXoadi(w http.ResponseWriter, fileName string) {

	pointerTem, _ := template.ParseFiles("./template/"+fileName, "./template/base.layout.tmpl")

	err := pointerTem.Execute(w, nil)
	if err != nil {
		fmt.Println("Error generating template")
		return
	}

}

// Problem, hiting disk is very cost. we should be hitting cache
var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, fileName string) {

	exitingValue := templateCache[fileName]
	if exitingValue != nil {
		fmt.Println("Hitting cache template")
		err := exitingValue.Execute(w, nil)
		if err != nil {
			fmt.Println("Error while generating template")
			return
		}
	} else {
		pointerTemplate, err := createTemplateCache(fileName)
		if err == nil {
			pointerTemplate.Execute(w, nil)
		}

	}

}

func createTemplateCache(fileName string) (*template.Template, error) {
	templates := []string{
		fmt.Sprintf("./template/%s", fileName),
		"./template/base.layout.tmpl",
	}

	pointerTemplate, err := template.ParseFiles(templates...)
	if err != nil {
		return nil, err
	}
	templateCache[fileName] = pointerTemplate
	return pointerTemplate, nil
}
