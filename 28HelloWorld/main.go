package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8000"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addvalues(2, 3)
	fmt.Fprintf(w, "this is about page")
	_, _ = fmt.Fprintf(w, "Addition of 2 & 3 is %d", sum)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	res, err := divide(2, 0)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
	_, _ = fmt.Fprintf(w, "division of 2 & 3 is %f", res)

}

func addvalues(x, y int) int {
	return x + y
}

func divide(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("divided by zero error")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/error", Divide)

	fmt.Printf("Server running on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
