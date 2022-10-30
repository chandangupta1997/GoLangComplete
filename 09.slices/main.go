package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(" welcoome to the slices class ")

	// array with lenght can be considered as slices

	var fruitList = []string{}

	fmt.Printf("type of frutList %T \n", fruitList)
	//type of frutList []string

	// here the length of array is not defined
	// so we can add things as much we want

	//fruitList[0] = "apple" //panic: runtime error: index out of range [0] with length 0

	// this is the only way to make slices

	fruitList = append(fruitList, "banana", "carrot")

	fmt.Println("fruitList:", fruitList)

	// lets make an array with help of make

	highScore := make([]int, 4)

	highScore[0] = 123
	highScore[1] = 31
	// length ke aage jaoege to error dega

	//but

	highScore = append(highScore, 3, 2, 1)

	fmt.Println("highScores:", highScore)

	// for sorting

	//sortedHighScore:= sort.Ints(highScore)
	sort.Ints(highScore)

	// for removing items from an array // slice

	var courses = []string{"reactJs", "javascript", "swift", "python", "ruby"}
	var indexToRemove int = 2
	courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...)

	fmt.Println(courses)

}
