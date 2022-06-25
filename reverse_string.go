package reverse_string

func ReverseString(input string) (output string) {
	inputSlice := []rune(input)
	outputSlice := make([]rune, len(inputSlice))
	for i := 0; i < len(inputSlice); i++ {
		outputSlice[i] = inputSlice[len(inputSlice)-1-i]
	}
	output = string(outputSlice)
	return output
}
