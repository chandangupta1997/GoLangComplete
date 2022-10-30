package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://youtube.com:3000/golang?time=10?&anotherQuery=ok?user=chandan"

func main() {

	fmt.Println(" welcome to the URL class")

	//package used  "net/url"

	result, _ := url.Parse(myUrl)
	// fmt.Println(result)

	// fmt.Print(result.Scheme)
	// fmt.Print(result.Host)
	// fmt.Print(result.Path)
	// fmt.Print(result.RawQuery)
	// fmt.Print(result.Port())

	queryParams := result.Query()

	fmt.Println(queryParams)
	fmt.Printf(" type of the qyeryParams %T\n", queryParams)

	//qyeryParamsurl.Values  this will be an object
	//so we can have a loop on that

	for _, val := range queryParams {

		fmt.Println("params  is ", val)

	}

	// now  reversing the task creting from chunks of data

	partsOfUrl := &url.URL{

		Scheme:   "https",
		Host:     "youtube.com",
		Path:     "watch",
		RawQuery: "v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26",
	}

	createdUrl := partsOfUrl.String()

	fmt.Println(createdUrl)

}
