package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/theseregrets/hotel-managment-system-go/pkg/config"
)

 var functions = template.FuncMap{}

 var app *config.AppConfig

 func NewTemplate(a *config.AppConfig)  {
	app=a
 }
 
// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc := app.TempCache
	
	t, ok := tc[tmpl]
	if !ok {
		fmt.Println("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)


	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}

}

// CreateTemplateCache creates a template cache as a map
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
