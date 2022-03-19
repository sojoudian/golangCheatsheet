package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	str := "Albus"
	fmt.Println("Type of str is: ", reflect.TypeOf(str))
	b, _ := strconv.ParseBool(str)
	fmt.Println("type of b", reflect.TypeOf(b))
	fmt.Println("Type of str is: ", reflect.TypeOf(str))

	flt, _ := strconv.ParseFloat(str, 64)
	fmt.Println("type of flt", reflect.TypeOf(flt))

	intg, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println("type of intg", reflect.TypeOf(intg))
}
