package strconv

func DeleteLeadingZeros(numberAsString []byte) []byte {
	for key, value := range numberAsString {
		if value != 48 {
			return numberAsString[key:]
		}
		if value == 48 && key+1 == len(numberAsString) {
			return []byte{48}
		}
	}
	return numberAsString
}
