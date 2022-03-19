package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5}
	j := []int{21, 22, 23, 24, 25, 26, 27, 28, 29}
	k := []int{31, 32, 33, 34, 35, 36, 37, 38, 39}

	fmt.Println("i: ", i, "and i capacity is: ", cap(i))
	fmt.Println("j: ", j, "and j capacity is: ", cap(j))
	fmt.Println("k: ", k, "and k capacity is: ", cap(k))

	sum := []int{}
	sum = append(i, j...)
	sum = append(sum, k...)
	fmt.Println("sum: ", sum, "and sum capacity is:", cap(sum))
}
