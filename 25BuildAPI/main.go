package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model

type User struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Task     *Task  `json:"task"`
}

type Task struct {
	TaskID    string    `json:"taskID"`
	TaskName  string    `json:"taskName"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

//DB
var users []User

func (u *User) IsEmpty() bool {
	// return u.UserID == "" && u.UserName == ""
	return u.UserName == ""
}

func main() {
	fmt.Println("User CRUD")

	r := mux.NewRouter()

	//seeding data
	users = append(users, User{UserID: "1", UserName: "kasi", Email: "kasi.srm@gmail.com",
		Task: &Task{TaskID: "a", TaskName: "Task 1", Status: true}})

	users = append(users, User{UserID: "2", UserName: "Magizhan", Email: "magizhan@gmail.com",
		Task: &Task{TaskID: "b", TaskName: "Task 2", Status: true}})

	//routing
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/users", GetAllUser).Methods("GET")
	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/user", AddOneUser).Methods("POST")
	r.HandleFunc("/user/{id}", UpdateOneUser).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteOneUser).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome API</h1>"))
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Users")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one user")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	fmt.Println(params)

	for _, user := range users {
		if user.UserID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode("No user found.")
	return
}

func AddOneUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add new user")
	w.Header().Set("Content-Type", "application/json")

	// check body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send input")
		return
	}
	// check body is {}
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.IsEmpty() {
		json.NewEncoder(w).Encode("Data is empty")
		return
	}

	// generate UID as string
	rand.Seed(time.Now().UnixNano())
	user.UserID = strconv.Itoa(rand.Intn(100))
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
	return
}

func UpdateOneUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one User")
	w.Header().Set("content-Type", "application/json")

	params := mux.Vars(r)
	for index, user := range users {
		fmt.Println(user.UserID)
		if user.UserID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.UserID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

func deleteOneUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one User")
	w.Header().Set("content-Type", "application/json")

	params := mux.Vars(r)

	for index, user := range users {
		if user.UserID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			json.NewEncoder(w).Encode("User deleted.")
			return
		}
	}
}
