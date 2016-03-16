package main

import (
	"classalike/user"
	"fmt"
)

func main() {
	const (
		timeFormat = "2006-01-02 15:00 UTC"
	)
	user := user.User{
		Firstname: "Pasit",
		Lastname:  "R.",
	}
	fmt.Printf("%p: %v\n", &user, &user)
	user.SetBirthdate(timeFormat, "1983-04-14 13:00 UTC")
	fmt.Printf("%p: %f\n", &user, user.Age())
}
