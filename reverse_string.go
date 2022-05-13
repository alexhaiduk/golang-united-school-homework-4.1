package reverse_string

func ReverseString(input string) (output string) {
	output = ""
	/*
		var temp = make([]rune, len(input))
		var i int = 0
		for _, v := range input {
			temp[i] = v
			i++
		}
		for i = len(temp) - 1; i >= 0; i-- {
			if temp[i] != 0 {
				output += string(temp[i])
			}
		}
		for i = 0; i < len(temp); i++ {
			if temp[i] != 0 {
				output = string(temp[i]) + output
			}
		}
	*/
	for _, v := range input {
		output = string(v) + output
	}
	return
}
