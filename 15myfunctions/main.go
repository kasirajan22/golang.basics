package main

import "fmt"

func main() {
	fmt.Println("Functions")

	resAdder := adder(3, 7)

	fmt.Println("Result adder is:", resAdder)

	resPro, msgPro := proAdder(5, 10)
	fmt.Println("Result from proAdder is:", resPro)

	fmt.Println("Message from proAdder is:", msgPro)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val

	}

	return total, "Message from proAdder"
}
