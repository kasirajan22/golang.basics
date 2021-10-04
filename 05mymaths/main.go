package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	fmt.Println("Welcome")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4.7

	// fmt.Println("the sum is: ", myNumberOne+int(myNumberTwo))

	//randomNumbers
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	//random from crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
