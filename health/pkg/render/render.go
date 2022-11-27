package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/fontexd/go/health/pkg/config"
	"github.com/fontexd/go/health/pkg/models"
)

var functions = template.FuncMap{}

var app *config.Appconfig

func NewTemplates(a *config.Appconfig) {
	app = a

}

func addDefaultData(td *models.Templatedata) *models.Templatedata {

	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.Templatedata) {
	var tc map[string]*template.Template
	// Create a template cache

	if app.UseCache {
		//Get the template cache from he app config
		tc = app.TemplateCache
	} else {

		tc, _ = CreateTemplateCache()
	}
	//get template from cache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get templatefrom cache")
	}

	//render the template

	buf := new(bytes.Buffer)

	td = addDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to Browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	//get all of th files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl*")
	if err != nil {
		return myCache, err
	}
	//range throug all files ending with *.page.tmpl

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
