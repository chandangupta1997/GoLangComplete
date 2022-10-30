package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(" welcome to our  app")
	fmt.Println(" please rate our pizza ")

	reader := bufio.NewReader(os.Stdin)

	// just like try catch

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating ", input)

	// data conversion kind of parsing

	//package used  strconv.ParseFloat()

	numRating, err := strconv.ParseFloat(input, 64)

	fmt.Println("hii there i am here ", numRating, err)

	//hii there i am here  0 strconv.ParseFloat: parsing "7\r\n": invalid syntax

	numRatings, err2 := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err2 != nil {
		fmt.Println(err2)
	} else {

		fmt.Println("we are adding 1 to your rating ", numRatings+1)

	}

}
