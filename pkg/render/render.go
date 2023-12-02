package render

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// Runing now
func RenderTemplate(w http.ResponseWriter, fileName string) {
	mapResult, err := createAllTemplate()
	if err != nil {
		log.Fatal("Template not found")
	}
	currentTemp := mapResult[fileName]
	currentTemp.Execute(w, nil)

}

// create a new tempalte that load all files in folder Template at once.
func createAllTemplate() (map[string]*template.Template, error) {

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

//var templateCache = make(map[string]*template.Template)

/*
// Render Template
func RenderTemplate_Pause(w http.ResponseWriter, fileName string) {
	existing := templateCache[fileName]
	if existing != nil {
		log.Println("existing  template of tc")
		err := existing.Execute(w, nil)
		if err != nil {
			fmt.Println("Error generating existing")
			return
		}
	} else {
		log.Println("create new template for tc")
		pTemplate := createNewTemplate(fileName)
		log.Println("after creating a new template for tc")

		err := pTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("Error generating pTemplate")
			return
		}
	}
}

// create a new template and store it in map cache
func createNewTemplate(fileName string) *template.Template {

	sliceFiles := []string{
		fmt.Sprintf("./template/%s", fileName),
		"./template/base.layout.tmpl",
	}

	pTemplate, err := template.ParseFiles(sliceFiles...)
	if err != nil {
		log.Println("create new template error")
		return nil
	} else {
		log.Println("add new  template to caches and return")
		templateCache[fileName] = pTemplate
		return pTemplate
	}
}

*/
