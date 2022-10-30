package main

import "fmt"

func main() {

	fmt.Println(" welcome to pointer class ")

	var pointer1 int

	var pointer2 string

	fmt.Println("value of pointer1 ", pointer1) // nil or 0
	fmt.Println("value of pointer2 ", pointer2) // kuch dikha hi nhi rha

	// second way of declaring
	a := 10
	myPointer1 := &a

	myPointer2 := &myPointer1

	fmt.Println("value/address of myPointer1", myPointer1) //value of b 0xc000016078

	// it is showing the address of the varibale not the value
	// for value

	fmt.Println(" only value of myPointer1", *myPointer1)

	fmt.Println("value of myPointer2//adress ", myPointer2) //value of myPointer2 0xc0000ca020
	fmt.Println("value of myPointer2 ", *myPointer2)

}
