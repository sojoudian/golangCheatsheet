package main

import "fmt"

type X struct {
	a string
	b int
}

func (varx X) methodX(n string) { // a method with value receiver
	varx.a = n
}

func (vary *X) methodY(num int) { // a method with pointer receiver} vary *Y is pointer receiver
	vary.b = num
}

func (varZ *X) methodZ(z string) { // a method with pointer receiver} vary *Y is pointer receiver
	varZ.a = z
}

func main() {
	m := X{
		"abc",
		123,
	}

	fmt.Printf("m.a = %s, m.b=%d\n", m.a, m.b)

	m.methodX("xyz") //it is not impacted the methodX , because it will create a
	//copy of varx.a , so when we made a modification like this the original will not get modified
	//but the coppied version will be change
	fmt.Printf("m.a = %s \n", m.a)
	m.methodY(1001)
	fmt.Printf("m.b = %d\n", m.b)
	(&m).methodY(1234)
	fmt.Printf("m.b = %d\n", m.b)
	fmt.Printf("m.a = %s, m.b=%d\n", m.a, m.b)
	(&m).methodZ("hey")
	fmt.Printf("m.a = %s, m.b=%d\n", m.a, m.b)
	//as we can see using the pointer receiver we can change the original string
}
