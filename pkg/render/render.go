package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	templateCache, err := CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	template, _ := templateCache[tmpl]

	buffer := new(bytes.Buffer)

	err = template.Execute(buffer, nil)

	if err != nil {
		log.Println(err)
	}

	_, err = buffer.WriteTo(w)

	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		pageName := filepath.Base(page)

		templateSet, err := template.New(pageName).ParseFiles(page)

		if err != nil {
			return templateCache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return templateCache, err
		}

		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return templateCache, err
			}
		}

		templateCache[pageName] = templateSet
	}

	return templateCache, nil
}
