package strconv

func DeleteLeadingZeros[I input](input I) I {
	for key := 0; key < len(input); key++ {
		if input[key] != 48 {
			return input[key:]
		}
		if input[key] == 48 && key+1 == len(input) {
			return input[key:]
		}
	}
	return input
}
