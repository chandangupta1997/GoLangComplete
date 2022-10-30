package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println(" welco9me to the  file classes ")

	// like any other languages you can only CRUD txt files for other you nedd libraries

	content := "i quit teaching today "

	// creating file using os package

	file, err := os.Create("./myFirstFile.txt")

	if err != nil {
		panic(err)
	}

	// writing content into the file using io

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("the length of content is ", length) //34
	defer file.Close()

	readFile("./myFirstFile.txt")

}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)

	if err != nil {

		panic(err)
	}
	fmt.Println(" file Content:", string(dataByte))

}

// func deleteFile(fileName string) {

// }
