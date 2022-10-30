package main

import "fmt"

func main() {

	fmt.Println(" welcome to the arrays ")

	// declaring an array

	var array [5]string

	fmt.Println(array) //[    ]

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"

	fruitList[3] = "baby"

	fmt.Println("here is your fruitList:", fruitList)

	//here is your fruitList: [apple banana  baby]
	// extraa space respresenting 3rd place is vacant

	var vegyList = [4]string{"a", "b", "c", "d"}

	fmt.Println("Here is your VegyList", vegyList)

	fmt.Println(len(array), len(fruitList))

	// len will show the lenght of array at the time of initialisation

}
