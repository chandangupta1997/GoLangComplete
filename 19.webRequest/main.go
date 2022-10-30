package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev/"

func main() {

	fmt.Println(" welcome to the go lang web request handler classes ")

	//https://pkg.go.dev/net/http

	//******************************** step 1 *********************************88888

	response, err := http.Get(url)

	//checking the type of the response

	fmt.Printf("the type of the response %T\n", response)
	//*http.Response  and this is a pointer maker sure that your are getting original copy of the data

	if err != nil {
		panic(err)
	}

	// it is caller's responsibility to close the connection
	defer response.Body.Close()

	//defer :=  used to delay the execution of a function or a statement until the nearby function returns

	// ******************************* step 2 *****************************************
	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	//******************************** step 3 *****************************************
	// stringify the data

	stringifiedContent := string(dataBytes)

	fmt.Println("this is the  stringifiedContent:", stringifiedContent)

}
