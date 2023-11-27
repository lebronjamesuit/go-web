package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, fileName string) {

	pointerTem, _ := template.ParseFiles("./template/" + fileName)

	err := pointerTem.Execute(w, nil)
	if err != nil {
		fmt.Println("Error generating template")
		return
	}

}
