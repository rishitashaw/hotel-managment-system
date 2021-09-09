package main

import (
	"fmt"
	"html/template"
	"net/http"
)


func renderTemplates(w http.ResponseWriter, tmplName string){
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmplName)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
