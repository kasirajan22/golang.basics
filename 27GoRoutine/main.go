package main

import "time"

func main() {
	go print("Hello")
	print("world")
}

func print(s string) {
	for i := 0; i < 4; i++ {
		time.Sleep(2 * time.Second)
		println(s)
	}
}
