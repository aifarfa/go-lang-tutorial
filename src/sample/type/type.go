package main

import (
	"fmt"
	// "log"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	user := CreateUser()
	fmt.Println(user)
	u := &user //pointer to address of user
	fmt.Printf("pointer: %p, value: %v\n", u, u)
}

func CreateUser() User {
	user := User{
		Age:       32,
		FirstName: "Pasit",
		LastName:  "R.",
	}
	return user
}
