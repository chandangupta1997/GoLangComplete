package main

import "fmt"

func main() {
	fmt.Println(" welcome to the loop classes ")

	// keyword only for

	//1 easy and familiar method

	arr := []string{"sunday", "monday", "tuesday", "wednesday", "thursday"}

	for i := 0; i < len(arr); i++ {

		fmt.Println(arr[i])

	}

	//second variation
	// i stand for index only

	for i := range arr {
		fmt.Println(arr[i])
	}

	// very immportant
	// kind of for Each
	// but here you can acces both index and values

	for index, value := range arr {
		fmt.Println(index, value)

	}

	// While loop

	num := 1

	for num < 10 {
		fmt.Println(" value of num is:", num)
		num++

		if num > 9 {
			goto lco

		}
	}

	// break continue and goto

lco:
	fmt.Println("this is the last print")

}
