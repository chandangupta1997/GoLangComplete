package main

import "fmt"

func main() {

	fmt.Println("welcome to the functions ")
	// this will be the entry point
	// you dont need to call this function

	adderResult := adder(2, 3)
	fmt.Println("the adder result is ", adderResult)

	proaAdderResult, argumentlength := proAdder(1, 2, 3, 4, 5)

	fmt.Println(" argumetn length is  ", argumentlength)
	fmt.Println(" the pro adder result is  which can accept n arguments ", proaAdderResult)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, int) {

	// just like spread operatoe
	// now you can accept n number of argument
	// it will an slice or array

	total := 0
	argumentLength := 0
	for i, val := range values {

		total = total + val
		argumentLength = i + 1
	}

	return total, argumentLength
}
