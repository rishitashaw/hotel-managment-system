package main

import (
	"fmt"
	"net/http"

	"github.com/theseregrets/hotel-managment-system-go/pkg/config"
	"github.com/theseregrets/hotel-managment-system-go/pkg/handlers"
	"github.com/theseregrets/hotel-managment-system-go/pkg/render"
)

const portNumber="8080"


func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()

	if err != nil {
		fmt.Println("cannot create template cache")
	}

	app.TempCache = tc
	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("started listening at http://localhost:%s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

	
}
