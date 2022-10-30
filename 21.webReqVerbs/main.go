// https://www.youtube.com/watch?v=V-sxFQ0fWlw&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=28
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println(" welcome to the webRequest Class")
	PerformGetRequest()
	PerformJsonRequest()
	PerformformRequest()

}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	// simple hitting web request
	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	// close the request is our responsibility

	defer response.Body.Close()

	fmt.Println(response.Body)
	//&{0xc00021a000 {0 0} false <nil> 0xbe6ac0 0xbe6bc0}
	// this is a pointer

	//checking thr type of

	fmt.Printf(" the type of response body %T\n", response.Body) //*http.bodyEOFSignal

	//to read it we need to use ioutil library

	contentInByte, _ := ioutil.ReadAll(response.Body)

	fmt.Println(contentInByte)
	// previously in 20.webRequest i named content as dataBytes
	//[123 34 109 101 115 115 97 103 101 34 58 34 72 101 108 108 111 32 102 114 111 109 32 108 101 97 114 110 67 111 100 101 111 110 108 105 110 101 46 105 110 34 125]

	fmt.Printf(" the type of content  %T\n", contentInByte) // []uint8

	// so to convert it into human readable manner we need to stringify this

	// **************************special step 3*****************************

	//fmt.Println(string(content)) //{"message":"Hello from learnCodeonline.in"}

	var responseString strings.Builder

	byteCount, _ := responseString.Write(contentInByte)

	fmt.Println(byteCount) // 43

	fmt.Println(responseString.String())
	//fmt.Println(responseString.Write(contentInByte))

}

func PerformJsonRequest() {

	const myPostUrl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
	 {
       "name":"chandan Gupta"
	   "sex":"male"

	 }
	 `)

	// responseBody will get a pointer i.e memory address
	// as server is sending the same body
	responseBody, err := http.Post(myPostUrl, "application/json", requestBody)

	fmt.Println(responseBody)
	if err != nil {
		panic(err)
	}

	defer responseBody.Body.Close()

	content, err := ioutil.ReadAll(requestBody)

	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	fmt.Println(" here is the post response")

	fmt.Println(string(content))

}

func PerformformRequest() {

	const myUrl = "http://localhost:8000/postform"

	// creating form data

	// url is a package
	formData := url.Values{}

	formData.Add("firstName", "mittal")
	formData.Add("lastName", "nothing")
	//
	//
	// further can be added in the same way

	response, err := http.PostForm(myUrl, formData)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// response wiil be a pointer

	content, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
