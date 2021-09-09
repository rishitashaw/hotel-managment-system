package main

import (
	"net/http"
)


func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.page.html")

}

