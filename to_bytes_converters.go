package faststrconv

const (
	uint10    = 10
	uint100   = 100
	uint1000  = 1000
	uint10000 = 10000
)

// Uint162Bytes converts an uint16 number to string.
func Uint162Bytes(num uint16) []byte {

	convertedNumber, i := getSliceUint16(num)

	for {
		convertedNumber[i] = byte(num%10) | 0x30
		num = num / 10
		if i == 0 {
			return convertedNumber
		}
		i--
	}
}

func getSliceUint16(num uint16) ([]byte, int) {
	if num < uint10 {
		return make([]byte, 1), 0
	}

	if num < uint100 {
		return make([]byte, 2), 1
	}

	if num < uint1000 {
		return make([]byte, 3), 2
	}

	if num < uint10000 {
		return make([]byte, 4), 3
	}

	return make([]byte, 5), 4
}

// Byte2Bytes converts a byte number to []byte.
func Byte2Bytes(num byte) []byte {
	convertedNumber, i := getSliceByte(num)

	for {
		convertedNumber[i] = num%10 | 0x30
		num = num / 10
		if i == 0 {
			return convertedNumber
		}
		i--
	}
}

func getSliceByte(num byte) ([]byte, int) {
	if num < uint10 {
		return make([]byte, 1), 0
	}

	if num < uint100 {
		return make([]byte, 2), 1
	}

	return make([]byte, 3), 2
}
