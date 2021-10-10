package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Make Post JSON request")
	performPostJSONRequest()
	fmt.Println("Make Post Form request")
	performPostFORMRequest()
}

func performPostJSONRequest() {
	myURL := "https://localhost:5001/post"

	//mock JSON Data

	requestBody := strings.NewReader(`
	{
		"name":"test"
	}
	`)

	response, err := http.Post(myURL, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCounts is: ", byteCount)
	fmt.Println(responseString.String())
}

func performPostFORMRequest() {
	myURL := "https://localhost:5001/postformencoder"

	//mock JSON Data

	data := url.Values{}
	data.Add("name", "Check")

	response, err := http.PostForm(myURL, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCounts is: ", byteCount)
	fmt.Println(responseString.String())
}
