package main

import (
	"fmt"
	"net/http"

	"github.com/hakem26/hello_web_go/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("%s", fmt.Sprintf("The Server Started On Port %s \n", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
