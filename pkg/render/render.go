package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// var functions = template.FuncMap{}

func RenderTemplates(w http.ResponseWriter, htmlName string){

	parsedTemplate,_ := template.ParseFiles("./templates/"+htmlName)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error in rendering templates", err)
		return
	}

	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)	
	// }

	// t, ok := tc[htmlName]
	// if !ok {
	// 	log.Fatal(err)
	// }

	// buf := new(bytes.Buffer)

	// _ = t.Execute(buf, nil)

	// _, err = buf.WriteTo(w)
	
	// if err != nil {
	// 	fmt.Println("error in getting template cache", err)
	// }

}

// func CreateTemplateCache() (map[string]*template.Template, error) {

// 	myCache := map[string]*template.Template{}

// 	pages, err := filepath.Glob("../templates/*.page.html")
// 	if err != nil {
// 		return myCache, err
// 	}

// 	for _, page := range pages {
// 		name := filepath.Base(page)
// 		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
// 		if err != nil {
// 			return myCache, err
// 		}

// 		matches, err := filepath.Glob("./templates/*.layout.html")
// 		if err != nil {
// 			return myCache, err
// 		}

// 		if len(matches) > 0 {
// 			ts, err = ts.ParseGlob("./templates/*.layout.html")
// 			if err != nil {
// 				return myCache, err
// 			}
// 		}

// 		myCache[name] = ts
// 	}

// 	return myCache, nil
// }
