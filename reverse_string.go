package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	runes := []rune(input)
	for i, j := 0, len(runes); i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return output
}
