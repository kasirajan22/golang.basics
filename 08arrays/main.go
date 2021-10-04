package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	//fruitList[2] = "Apple"
	fruitList[3] = "orange"

	fmt.Println("Fruit List is:", fruitList)
	fmt.Println("Fruit List is:", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Mushroom"}

	fmt.Println("vegy list is:", len(vegList))
}
