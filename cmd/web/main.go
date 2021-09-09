package main

import (
	"fmt"
	"net/http"

	"github.com/theseregrets/hotel-managment-system-go/pkg/handlers"
)

const portNumber="8080"


func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("started listening at http://localhost:%s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
