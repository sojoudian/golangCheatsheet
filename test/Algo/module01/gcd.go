package module01

// func GCD(a, b int) int {
// 	for b != 0 {
// 		a, b = b, a%b
// 	}
// 	return a
// }

// func GCD(a, b int) int {
// 	// step 1: if  b == 0, return A
// 	if b == 0 {
// 		return a
// 	}
// 	//Step 2: A becomes B, and B becomes the remainder of division A by B
// 	// `a, b = b, a%b`
// 	a, b = b, a%b
// 	return GCD(a, b)
// }

// func GCD(a, b int) int {
// 	for {
// 		if b == 0 {
// 			return a
// 		}
// 		a, b = b, a%b
// 		return GCD(a, b)
// 	}
// }

func GCD(a, b int) int {
	for b != 0 {
		if b == 0 {
			return a
		}
		a, b = b, a%b
	}
	return a
}
