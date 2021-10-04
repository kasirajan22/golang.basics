package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")

	var ptrOne *int
	fmt.Println("the pointer is :", ptrOne)

	mynumber := 23

	var ptr = &mynumber
	fmt.Println("value of pointer is: ", ptr)

	fmt.Println("value of pointer is: ", *ptr)

	*ptr = *ptr + 1

	fmt.Println("New Value: ", mynumber)
}
