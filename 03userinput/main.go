package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating")

	input, _ := reader.ReadString('\n')
	fmt.Println("Rating is ", input)

	fmt.Printf("rating type is: %T \n", input)
}
