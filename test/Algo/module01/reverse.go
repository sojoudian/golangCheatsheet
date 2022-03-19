package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var reversed string
	// for i := len(word) - 1; i >= 0; i-- {
	// 	reversed = reversed + string(word[i])
	// }

	// for i := 0; i < len(word); i++ {
	// 	reversed = string(word[i]) + reversed
	// }

	for _, v := range word {
		reversed = string(v) + reversed
	}
	return reversed

}
