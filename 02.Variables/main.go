package main

import "fmt"

const LoginToken = "MyToken" // L  capital L showing that it is a public

func main() {

	// String

	var username string = "Chandan"
	fmt.Println(username)

	fmt.Printf(" varibale type of :%T \n ", username)

	// Integer

	var num int = 1233456
	fmt.Println(num)
	fmt.Printf(" varibale type of :%T \n ", num)

	//Float

	var floatnumber float64 = 12345.123

	fmt.Println(floatnumber)
	fmt.Printf(" varibale type of :%T \n ", floatnumber)

	// Boolean

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)

	fmt.Printf(" varibale type of :%T \n ", isLoggedIn)

	// default value and aliases

	// only declaration no assignment   in Js we will get undefined

	var novalue int // will get 0
	fmt.Println(novalue)
	fmt.Printf(" varibale type of :%T \n ", novalue) // type will be  int

	var novalueString string
	fmt.Println(novalueString)                             // no value  is assigned
	fmt.Printf(" varibale type of :%T \n ", novalueString) // string

	// another ways
	// implicit type
	var website = "learncode"
	fmt.Println(website)
	fmt.Printf(" varibale type of :%T \n ", website)

	// no var style

	numberOfUser := 300000 // only allowed inside a method
	fmt.Println(numberOfUser)
	fmt.Printf(" varibale type of :%T \n ", numberOfUser)

	fmt.Println(LoginToken) // from public

}
