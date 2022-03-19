package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 1; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

}

// I like this commendted version more
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	for i := 65; i <= 90; i++ {
// 		fmt.Printf("%v\t%#U\n", i, i)

// 	}

// }
