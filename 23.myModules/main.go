package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println(" welcome to the mode classes ")
	greeter()
	r := mux.NewRouter() // router module
	r.HandleFunc("/", serveHome)
	log.Fatal(http.ListenAndServe(":4040", r))

}

func greeter() {
	fmt.Println("aur bhai sahb ki haal chaal ")
}

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> this is the home page</h1>"))

}
