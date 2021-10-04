package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files")
	content := "Content to write in file"

	file, err := os.Create("./myfile.txt")
	showError(err)

	length, err := io.WriteString(file, content)
	showError(err)
	fmt.Println("lenght is, ", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	showError(err)

	fmt.Println("Value in the file is, \n", databyte)
	fmt.Println("Value in the file is,\n", string(databyte))
}

func showError(err error) {
	if err != nil {
		panic(err)
	}
}
