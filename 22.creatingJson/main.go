package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	// aliases
	Name      string `json:"courseName"`
	Price     int
	Platforms string   `json:"websites"`
	Tags      []string `json:"tags,omitempty"` // if fiels is empty this field will not show
	password  string   `json:"-"`              // if you dont want to show this field in the response
}

func main() {
	fmt.Println("welcome to the creating json class")

	//encodeJson()
	decodedJson()

}

func encodeJson() {
	myCourses := []course{
		{"React", 299, "udemy.com", []string{"js,React"}, "abc123"},
		{"DSA", 199, "youtube.com", []string{"DSA"}, "123abc"},
		{"GO", 0, "youtube.com", nil, "145Qa"},
	}

	// marshal return json encoding format
	finalJson, err := json.Marshal(myCourses)
	if err != nil {
		panic(err)
	}

	fmt.Println(finalJson) // bytecode json

	// in both of the case we can read the stuff
	fmt.Printf("%s\n", finalJson)

	fmt.Println(string(finalJson))

	// indentet  json

	indententJson, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(indententJson))

}

func decodedJson() {

	// everything coming from web will be in byte format
	// so for byte to json we will so the following
	jsonFromWeb := []byte(`{
		"courseName":"React",
		"Price": 299,
		"websites": "udemy.com",
		"tags": ["js,React"]
		}`)

	//checking valid json or not

	var myCourses course

	isValidJson := json.Valid(jsonFromWeb)

	if isValidJson {
		fmt.Println("valid json ")
		json.Unmarshal(jsonFromWeb, &myCourses) // decode according to the strucutre (& and for reference )

		fmt.Printf("%#v\n", myCourses)

	} else {
		fmt.Println("not valid json ")
	}

}
