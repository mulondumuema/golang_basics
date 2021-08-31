package main

import (
	"fmt"
	"log"
	"net/http"
)

// in web, we are sending a request and waiting for a response
// because we are building an application which is not the web browser but the web server, we receive the request and we build the response
func homePage(w http.ResponseWriter, r *http.Request) { // this funtion is a handler
	html := `<strong>Hello, world</strong>`
	w.Header().Set("Content-Type", "text/html") // tells the browser that you are about to receive some html content
	fmt.Fprint(w, html) // this prints the response on the browser
}

func main() {
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil) // this starts a production ready web server and listens on port 8080 
}