package main

import "fmt"

type User struct {
	Name, Role, Email string
	// Name  string
	// Role  string
	// Email string
	Age int
}

func (u User) Salary() int {
	switch u.Role {
	case "Student":
		return 10
	case "Headmaster":
		return 100
	}
	return 0
}

func main() {
	tom := User{Name: "Tom", Role: "Student", Email: "tom@pois.hg.wiz", Age: 16}
	alb := User{Name: "Albus", Role: "Headmaster", Email: "albus@hg.wiz", Age: 116}

	fmt.Println(tom.Name, tom.Salary())
	fmt.Println(alb.Name, alb.Salary())
}
