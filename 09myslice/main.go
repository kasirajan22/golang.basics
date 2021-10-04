package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "banana", "Peach"}
	fmt.Printf("type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Tomato")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 335
	highScores[2] = 237
	highScores[3] = 239
	//highScores[4] = 333

	highScores = append(highScores, 444, 555, 666)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"Dotnet", "java", "Swift", "Golang", "Javascript"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
