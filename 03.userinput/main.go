package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("hey there welcome to my Pizza app  ")
	fmt.Println(" please rate our pizza  ")

	reader := bufio.NewReader(os.Stdin) // this line  created a buffer reader.

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating ", input)

}
