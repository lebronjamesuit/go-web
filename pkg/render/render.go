package render

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"jamesvo.uk/website/pkg/config"
)

// Make template caches accessible to AppConfig

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

// Runing now
func RenderTemplate(w http.ResponseWriter, fileName string) {

	mapResult := app.MyCaches
	currentTemp := mapResult[fileName]

	//
	log.Println("File name is" + fileName)
	log.Println("Current temp " + currentTemp.Name())
	currentTemp.Execute(w, nil)
}

// create a new tempalte that load all files in folder Template at once.
func CreateTemplateCaches() (map[string]*template.Template, error) {

	log.Println("render pgk, file templateCaches.go, func  CreateTemplateCaches")

	myCaches := make(map[string]*template.Template)
	sliceOfPathFilesName, err := filepath.Glob("./template/*.page.tmpl")
	if err != nil {
		log.Fatal("Error when load template", err)
		return nil, err
	}

	for _, path := range sliceOfPathFilesName {
		fileName := filepath.Base(path)
		ts, err := template.ParseFiles(path)
		if err != nil {
			log.Fatal("error creating Template")
		}
		layout := "./template/*.layout.tmpl"
		matches, err1 := filepath.Glob(layout)
		if err1 == nil && len(matches) > 0 {
			ts, _ = ts.ParseGlob(layout) // Attach layout to page
		}

		myCaches[fileName] = ts // template
	}

	// return the map
	return myCaches, nil

}
