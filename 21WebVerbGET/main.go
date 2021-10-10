package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Perform Get Request")
	performGetRequest()
}

func performGetRequest() {
	const myURL = "https://localhost:5001/weatherforecast"

	response, err := http.Get(myURL)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code is: ", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(content))

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCounts is: ", byteCount)
	fmt.Println(responseString.String())
}
