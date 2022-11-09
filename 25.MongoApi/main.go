package main

import (
	"fmt"
	"mongoapi/router"
	"net/http"
)

func main() {

	fmt.Println("welcome to the main function home route ")

	r := router.Router()

	http.ListenAndServe("2000", r)

}
