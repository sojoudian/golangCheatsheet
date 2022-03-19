package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name, Role, Email string
	// Name  string
	// Role  string
	// Email string
	Age int
}

func (u User) Salary() (int, error) {
	if u.Role == "" {
		return 0, errors.New(" Empty Role is not allowed")
	}
	switch u.Role {
	case "Student":
		return 10, nil
	case "Headmaster":
		return 100, nil
	}
	return 0, errors.New(
		fmt.Sprintf("I'm not able to handle the '%s' role", u.Role),
	)
}

func main() {
	// user := User{Name: "Albus", Role: "Headmaster", Email: "albus@hg.wiz", Age: 116}
	user := User{Name: "Albus", Role: "Unknown", Email: "albus@hg.wiz", Age: 116}
	//salary, err := user.Salary() it is equal to the code below
	if salary, err := user.Salary(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Salary: ", salary)
	}

}
