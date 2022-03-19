package main

import "fmt"

func main() {
	var input, rx, temp int
	var rev int = 0
	fmt.Println("Enter the number")
	fmt.Scanln(&input)

	temp = input
	for {
		rx = input % 10   //rx for 121 is 1
		rev = rev*10 + rx // rev is 0 and ===> (0*10)+1=1
		fmt.Println(rev)  //1 will be printed
		input /= 10
		if input == 0 {
			break
		}
	}
	if temp == rev {
		fmt.Println("The entered number is a palindromic number", temp)
	} else {
		fmt.Println("The entered number is not a palindromic number", rev)
	}
}
