package module01

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	// res := 0
	// multiplier := 1
	// l := len(value) - 1
	// for i := l; i >= 0; i-- {
	// 	var val int
	// 	fmt.Sscanf(string(value[i]), "%X", &val)
	// 	res = res + (multiplier * val)
	// 	multiplier = multiplier * base
	// }
	// return res

	const charset = "0123456789ABCDEF"
	multiplier := 1
	res := 0
	l := len(value) - 1
	for i := l; i >= 0; i-- {
		index := -1
		for j, char := range charset {
			if char == rune(value[i]) {
				index = j
				break
			}
		}
		res = res + (index * multiplier)
		multiplier = multiplier * base
	}
	return res
}
