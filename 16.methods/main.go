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
	user1.GetStatus()
	user1.ChangeEmail()

}

// defining structure here

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}

// creating methods here

func (u User) GetStatus() {

	if u.status {

		fmt.Print("status of user1 is active")

	} else {
		fmt.Println("user not active")
	}

}

func (u User) ChangeEmail() {
	u.Email = "changedemail@gmail.com"

	fmt.Println("changed email of user", u.Email)

	// but a copy is passed when a method is called on it

	// so changed happens in copy

	// to change in original  we will use pointers   i.e  *u

}
