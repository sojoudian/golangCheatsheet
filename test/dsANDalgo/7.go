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

func (u User) UpdatedNotWorking(email string) {
	u.Email = email
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
}

func main() {
	// tom := User{Name: "Tom", Role: "Student", Email: "tom@pois.hg.wiz", Age: 16}
	// alb := User{Name: "Albus", Role: "Headmaster", Email: "albus@hg.wiz", Age: 116}

	// fmt.Println(tom.Name, tom.Salary())
	// fmt.Println(alb.Name, alb.Salary())
	user := User{Name: "Tom", Role: "Student", Email: "tom@pois.hg.wiz", Age: 16}
	fmt.Println("Original user data struct", user)
	user.UpdatedNotWorking("a.dumbledore@hg.wiz")
	fmt.Println("After receiver method without pointer, (it is not working)", user)
	fmt.Println("---------------------------------------")
	user.UpdateEmail("a.dumbledore@hg.wiz")
	fmt.Println("Now Email is updated successfully with Pointer to *User", user)
}
