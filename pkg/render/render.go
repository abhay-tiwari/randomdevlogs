package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/abhay-tiwari/randomdevlogs/pkg/config"
	"github.com/abhay-tiwari/randomdevlogs/pkg/models"
	"github.com/justinas/nosurf"
)

var appConfig *config.AppConfig

func NewTemplates(config *config.AppConfig) {
	appConfig = config
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	var templateCache map[string]*template.Template

	if appConfig.UseCache {
		templateCache = appConfig.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	template, _ := templateCache[tmpl]

	buffer := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err := template.Execute(buffer, td)

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
			log.Println(err)
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
