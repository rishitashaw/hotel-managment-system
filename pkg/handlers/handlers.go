package handlers

import (
	"net/http"

	"github.com/theseregrets/hotel-managment-system-go/pkg/render"
)


func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.tmpl")
	
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.tmpl")

}

