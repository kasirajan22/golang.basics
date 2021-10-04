package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in golang")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01/02/2006 15:05:06 Monday"))

	createdDate := time.Date(2020, time.August, 7, 13, 10, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("01 02 2006 Monday"))
}
