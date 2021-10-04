package main

import "fmt"

func main() {
	fmt.Println("maps")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["GO"] = "golang"

	fmt.Println("list of languages:", languages)

	fmt.Println("Go shorts for:", languages["GO"])

	delete(languages, "RB")

	fmt.Println("List of languages is:", languages)

	// loops
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
