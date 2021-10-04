package main

import "fmt"

func main() {
	var username string = "kasi"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var anothervar string
	fmt.Println(anothervar)
	fmt.Printf("variable is of type: %T \n", anothervar)
}
