package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://pkg.go.dev/"

func main() {
	fmt.Println("Working with web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("respone type is: %T\n", response)

	defer response.Body.Close() // user must close.

	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)

	fmt.Println(content)
}
