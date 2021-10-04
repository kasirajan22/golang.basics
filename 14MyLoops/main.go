package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	rougValue := 1

	for rougValue < 10 {

		// if rougValue == 5 {
		// 	break
		// }
		if rougValue == 2 {
			goto kasi
		}
		if rougValue == 5 {
			rougValue++
			continue
		}

		fmt.Println("value is:", rougValue)
		rougValue++
	}

kasi:
	fmt.Println("Am Kasi")
}
