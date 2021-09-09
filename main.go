package main

import (
	"fmt"
	"net/http"
)

const portNumber="8080"


func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("started listening at http://localhost:%s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
