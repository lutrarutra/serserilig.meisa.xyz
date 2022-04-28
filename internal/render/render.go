package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/iMeisa/serserilig.meisa.xyz/internal/config"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/funcs"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/models"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/network"
	"github.com/justinas/nosurf"
)

var functions = funcs.Functions

var app *config.AppConfig

// NewRenderer sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds data for all templates
func AddDefaultData(data *models.TemplateData, r *http.Request) *models.TemplateData {
	stringMap := make(map[string]string)
	if data.StringMap != nil {
		stringMap = data.StringMap
	}

	stringMap["remote_ip"] = network.GetRealIP(r)

	data.StringMap = stringMap
	data.CSRFToken = nosurf.Token(r)
	return data
}

func Template(w http.ResponseWriter, r *http.Request, tmpl string, data *models.TemplateData) {

	var templateCache map[string]*template.Template

	if app.UseCache {
		// Get the page cache from the app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	page, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("page is not ok")
	}

	buf := new(bytes.Buffer)

	data = AddDefaultData(data, r)

	_ = page.Execute(buf, data)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing page to browser:", err)
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
