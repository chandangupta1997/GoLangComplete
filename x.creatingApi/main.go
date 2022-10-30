package main

import (
	"fmt"
	"log"
	"net/http"
)

// home page function
// handle request for our root URL

func homepage(w http.ResponseWriter, r *http.Request) {

	// printing response in browser
	fmt.Fprint(w, "homepage endpoint hit ")

}

func handleRequest() {

	http.HandleFunc("/", homepage)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {
	handleRequest()
}
