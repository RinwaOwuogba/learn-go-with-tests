package iteration

func Repeat(character string, noOfRepeats int) string {
	result := ""

	for i := 0; i < noOfRepeats; i++ {
		result += character
	}

	return result
}
