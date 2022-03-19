package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "Albus"
	b := 106
	c := 3.14
	d := true
	e := []string{"a", "b", "c", "d"}
	f := map[string]int{"first": 1, "second": 2, "third": 3}
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(e))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println("#######################################################")
	fmt.Println(reflect.ValueOf(a).Kind())
	fmt.Println(reflect.ValueOf(b).Kind())
	fmt.Println(reflect.ValueOf(c).Kind())
	fmt.Println(reflect.ValueOf(d).Kind())
	fmt.Println(reflect.ValueOf(e).Kind())
	fmt.Println(reflect.ValueOf(f).Kind())
	fmt.Println("#######################################################")
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)

}
