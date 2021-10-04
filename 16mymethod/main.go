package main

import (
	"fmt"
)

func main() {
	fmt.Println("Struct")
	// no inheritences. no parent class
	kasi := User{"kasi", "kasi@gmail.com", true, 30}

	fmt.Println(kasi)
	fmt.Printf("user detail: %+v\n", kasi)

	fmt.Printf("Name is %v and email is %v\n", kasi.Name, kasi.Email)

	kasi.GetStatus()

	kasi.NewMail()

	fmt.Printf("Name is %v and email is %v\n", kasi.Name, kasi.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (usr User) GetStatus() {
	fmt.Println("Is User active:", usr.Status)
}

func (usr User) NewMail() {
	usr.Email = "kasi@go.dev"
	fmt.Println("New Email is:", usr.Email)
}
