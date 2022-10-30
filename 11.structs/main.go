package main

import "fmt"

func main() {

	fmt.Println("welcome to the struct classes ")

	// we dont have classes in GO we have struct

	// there is no inheritance in GO No Parent No Child .

	user1 := User{"chandan", "chandanguptajobs@gmail.com", true, 25}

	fmt.Println(user1) //{chandan chandanguptajobs@gmail.com true 25}

	fmt.Printf("user1 details are :,%+v\n", user1)

	//user1 details are :,{Name:chandan Email:chandanguptajobs@gmail.com status:true Age:25}

}

// defining structure here

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}
