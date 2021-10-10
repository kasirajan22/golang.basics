package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://pkg.dev:3000/learn?coursename=Go&payment=4evdfbddfg"

func main() {
	fmt.Println("Working with URLs")
	fmt.Println(myURL)

	result, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	parms := result.Query()

	fmt.Printf("parameter type is %T\n", parms)

	// fmt.Println(parms["coursename"])
	// fmt.Println(parms["payment"])
	for _, val := range parms {
		fmt.Println("the parameter is: ", val)
	}

	//Creating URL

	newURL := &url.URL{
		Scheme:   "https",
		Host:     "pkg.dev",
		Path:     "learn",
		RawQuery: "user=kasi",
	}

	urlString := newURL.String()

	fmt.Println(urlString)
}
