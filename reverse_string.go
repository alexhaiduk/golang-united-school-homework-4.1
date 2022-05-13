package reverse_string

func ReverseString(input string) (output string) {
	output = ""
	for i := len(input) - 1; i >= 0; i-- {
		output = output + string(input[i])
	}
	return
}
