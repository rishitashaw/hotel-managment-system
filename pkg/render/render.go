package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}

func RenderTemplates(w http.ResponseWriter, tmplName string){

	_, err := RenderTemplatesTest(w)
	if err != nil {
		fmt.Println("error in getting template cache", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/"+tmplName)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

func RenderTemplatesTest(w http.ResponseWriter) (map[string]*template.Template, error){
	
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("page is currently on", page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches)>0{
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}