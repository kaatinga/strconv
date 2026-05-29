package faststrconv

const (
	uint10    = 10
	uint100   = 100
	uint1000  = 1000
	uint10000 = 10000
)

// Uint642Bytes converts an uint64 number to []byte.
func Uint642Bytes(num uint64) []byte {
	convertedNumber := make([]byte, decDigits64(num))
	i := len(convertedNumber) - 1
	for {
		convertedNumber[i] = byte(num%10) | 0x30
		num = num / 10
		if num == 0 {
			return convertedNumber
		}
		i--
	}
}

// Uint322Bytes converts an uint32 number to []byte.
func Uint322Bytes(num uint32) []byte {
	convertedNumber := make([]byte, decDigits32(num))
	i := len(convertedNumber) - 1
	for {
		convertedNumber[i] = byte(num%10) | 0x30
		num = num / 10
		if num == 0 {
			return convertedNumber
		}
		i--
	}
}

type unsigned interface {
	~uint | ~uint32 | ~uint64
}

// Uint2Bytes converts an uint, uint32 and uint64 number to []byte.
func Uint2Bytes[UI unsigned](num UI) []byte {
	convertedNumber := make([]byte, decDigits64(uint64(num)))
	i := len(convertedNumber) - 1
	for {
		convertedNumber[i] = byte(num%10) | 0x30
		num = num / 10
		if num == 0 {
			return convertedNumber
		}
		i--
	}
}

func decDigits32(num uint32) int {
	switch {
	case num < 10:
		return 1
	case num < 100:
		return 2
	case num < 1000:
		return 3
	case num < 10000:
		return 4
	case num < 100000:
		return 5
	case num < 1000000:
		return 6
	case num < 10000000:
		return 7
	case num < 100000000:
		return 8
	case num < 1000000000:
		return 9
	default:
		return 10
	}
}

func decDigits64(num uint64) int {
	switch {
	case num < 10:
		return 1
	case num < 100:
		return 2
	case num < 1000:
		return 3
	case num < 10000:
		return 4
	case num < 100000:
		return 5
	case num < 1000000:
		return 6
	case num < 10000000:
		return 7
	case num < 100000000:
		return 8
	case num < 1000000000:
		return 9
	case num < 10000000000:
		return 10
	case num < 100000000000:
		return 11
	case num < 1000000000000:
		return 12
	case num < 10000000000000:
		return 13
	case num < 100000000000000:
		return 14
	case num < 1000000000000000:
		return 15
	case num < 10000000000000000:
		return 16
	case num < 100000000000000000:
		return 17
	case num < 1000000000000000000:
		return 18
	case num < 10000000000000000000:
		return 19
	default:
		return 20
	}
}

// Uint162Bytes converts an uint16 number to []byte.
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
	const digit0 byte = '0'

	if num < uint10 {
		return []byte{num | digit0}
	}

	if num < uint100 {
		return []byte{
			(num / 10) | digit0,
			(num % 10) | digit0,
		}
	}

	return []byte{
		(num / 100) | digit0,
		((num / 10) % 10) | digit0,
		(num % 10) | digit0,
	}
}
