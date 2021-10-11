package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string `json:"User Name"`
	EMail    string
	Password string `json:"-"`
	Age      int
	Role     []string `json:"ROLE,omitempty"`
}

func main() {
	fmt.Println("Working with JSON")

	//EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {
	userDetail := []user{
		{"Test 1", "test@1.com", "test1234", 21, []string{"Admin", "User"}},
		{"Test 2", "test@2.com", "test1234", 22, []string{"Admin", "User", "Others"}},
		{"Test 3", "test@3.com", "test1234", 23, nil},
	}

	finalJSON, err := json.MarshalIndent(userDetail, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)
}

func DecodeJSON() {
	jsonFromWeb := []byte(`
	{
		"User Name": "Test 2",
		"EMail": "test@2.com",
		"Age": 22,
		"ROLE": ["Admin","User","Others"]
	}
	`)

	var data user

	if json.Valid(jsonFromWeb) {
		fmt.Println("Valid JSON")
		json.Unmarshal(jsonFromWeb, &data)
		fmt.Printf("%#v\n", data)
	} else {
		fmt.Println("InValid JSON")
	}

	var myUserData map[string]interface{}

	json.Unmarshal(jsonFromWeb, &myUserData)
	fmt.Printf("%#v\n", myUserData)

	for k, v := range myUserData {
		fmt.Printf("Key is %v and Value is %v and Type is %T\n", k, v, v)
	}
}
