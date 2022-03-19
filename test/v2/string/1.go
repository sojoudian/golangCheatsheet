package main

import "fmt"

func fn(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\t", s[i])
	}
	for i, v := range s {
		fmt.Println(i, v)
	}
}

func main() {
	s := "I am lord voldemort"
	b := "voldemort"
	fmt.Println(s)
	fn(b)

}
