package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"web-app/pkg/config"
	"web-app/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// SetConfig sets the config to a template package
func SetConfig(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate is rendering template without layout
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	tmpl = tmpl + ".page.tmpl"

	var tc map[string]*template.Template

	//Use cache or not checking
	if app.UseCache == true {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Template not found: ", tmpl)
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal("Error rendering template:", err)
	}
}

// CreateTemplateCache is rendering templates with layouts
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
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
	return myCache, nil
}
