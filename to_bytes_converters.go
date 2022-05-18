package strconv

const (
	uint1610    uint16 = 10
	uint16100   uint16 = 100
	uint161000  uint16 = 1000
	uint1610000 uint16 = 10000
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
	if num < uint1610 {
		return make([]byte, 1), 0
	}

	if num < uint16100 {
		return make([]byte, 2), 1
	}

	if num < uint161000 {
		return make([]byte, 3), 2
	}

	if num < uint1610000 {
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

const (
	byte10  byte = 10
	byte100 byte = 100
)

func getSliceByte(num byte) ([]byte, int) {
	if num < byte10 {
		return make([]byte, 1), 0
	}

	if num < byte100 {
		return make([]byte, 2), 1
	}

	return make([]byte, 3), 2
}
