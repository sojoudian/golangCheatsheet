package main

import (
	"fmt"
	"reflect"
)

func main() {
	ve(1, "a", true, 3.14, []string{"123", "string", "array"})
}

func ve(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, ":", reflect.ValueOf(v).Kind())
	}
}
