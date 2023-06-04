package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/whoajitpatil/reflecton/pkg/config"
	"github.com/whoajitpatil/reflecton/pkg/models"
)

var app *config.AppConfig

// NewTemplates : Sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData : Adds a default data for all handlers
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplates renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Coule not get the template from TC")
	}
	buf := new(bytes.Buffer)
	
	td = AddDefaultData(td)
	
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, err
}
