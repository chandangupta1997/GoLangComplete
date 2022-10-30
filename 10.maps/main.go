package main

import "fmt"

func main() {

	fmt.Println("welcome to the go maps ")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "python"

	fmt.Println("here is our map \n", languages)

	//here is our map
	//map[JS:javascript PY:python RB:Ruby]
	//like no commas in between

	// for accessing

	fmt.Println("for key JS value is ", languages["JS"]) //for key JS value is  javascript

	//for deleting

	// in js  and in Go

	delete(languages, "RB")

	fmt.Println(languages) //map[JS:javascript PY:python]

}
