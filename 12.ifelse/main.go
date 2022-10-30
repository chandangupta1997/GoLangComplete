package main

import "fmt"

func main() {

	fmt.Println("welcome to if else ")

	var voterAge int = 18

	if voterAge < 18 {
		fmt.Println("you can not  vote")
	} else {
		fmt.Println("you can  vote AChe din aaaenge ")
	}

	//special cases
	//when having value from the web request and checking or kind of validation

	num := 3
	if num%2 == 0 {
		fmt.Println("even number ")
	} else {

		fmt.Println(" odd number ")

	}

	// direct mathematical condition

	if false {

	} else if true {

	} else if true {

	} else {

	}

}
