package faststrconv

// DeleteLeadingZeros removes all the leading zeros in the input string.
func DeleteLeadingZeros[I stringOrBytes](input I) I {
	if len(input) == 0 {
		return input
	}

	key := 0
	for ; key < len(input); key++ {
		if input[key] != 48 {
			return input[key:]
		}
	}
	return input[key-1:]
}
