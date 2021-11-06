package main

import (
	"fmt"
	"log"
	"mangodbapi/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")

	r := router.Router()
	fmt.Println("Server is gettting started...")
	err := http.ListenAndServe(":4000", r)
	fmt.Println(err)
	log.Fatal(err)
	fmt.Println("Listening port 4000...")
}
