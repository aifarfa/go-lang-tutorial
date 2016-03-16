package user

import (
	"fmt"
	"time"
)

type User struct {
	Firstname, Lastname string
	Id                  int
	Birthdate           time.Time
}

func (u User) Fullname() string {
	return u.Firstname + " " + u.Lastname
}

func (u User) Age() int {
	duration := time.Since(u.Birthdate)
	current := time.Now().Year()
	birthYear := u.Birthdate.Year()

	fmt.Printf("age: %f hours, since: %d - %d\n", duration.Hours(), birthYear, current)
	return current - birthYear
	// return duration.Hours() / 24.0 / 365.0 // TODO
}

// retrieve pointer ref
func (u *User) SetBirthdate(timeFormat, value string) {
	date, err := time.Parse(timeFormat, value)
	if err != nil {
		fmt.Println(err)
	} else {
		u.Birthdate = date
	}
}
