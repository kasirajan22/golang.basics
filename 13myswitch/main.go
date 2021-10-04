package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")
		fallthrough
	case 2:
		fmt.Println("Dice valus is 2")
	case 3:
		fmt.Println("Dice valus is 3")
	case 4:
		fmt.Println("Dice valus is 4")
	case 5:
		fmt.Println("Dice valus is 5")
	case 6:
		fmt.Println("Dice valus is 6")
	default:
		fmt.Println("other than 6")
	}
}
