package strconv

type input interface {
	~[]byte | ~string
}

// GetByte validates and converts input string or []byte to byte type.
func GetByte[I input](input I) (byte, error) {

	if len(input)&^ByteLengthMask != 0 || len(input) == 0 {
		return 0, ErrNotByte
	}

	var i int
	var output uint16
	for {
		if input[i] < 0x30 || input[i] > 0x39 {
			return 0, ErrNotByte
		}

		output = uint16(input[i])&UnicodeMask + output*10

		i++

		if i == len(input) {
			if output&^0xff != 0 {
				return 0, ErrNotByte
			}
			break
		}
	}

	return byte(output), nil
}

// GetUint16 validates and converts input string or []byte to uint16 type.
func GetUint16[I input](input I) (uint16, error) {
	if len(input) == 0 {
		return 0, ErrNotUint16
	}

	var i int
	var output uint32
	for {

		if input[i] < 0x30 || input[i] > 0x39 {
			return 0, ErrNotUint16
		}

		output = output*10 + uint32(input[i])&UnicodeMask

		i++

		if i == len(input) {
			if output&^0xffff != 0 {
				return 0, ErrNotUint16
			}
			break
		}
	}

	return uint16(output), nil
}

// GetUint32 validates and converts input string or []byte to uint32 type.
func GetUint32[I input](input I) (uint32, error) {

	if len(input) == 0 {
		return 0, ErrNotUint32
	}

	var i int
	var output uint64
	for {

		if input[i] < 0x30 || input[i] > 0x39 {
			return 0, ErrNotUint32
		}

		output = uint64(input[i])&UnicodeMask + output*10

		i++

		if i == len(input) {
			if output&^0xffffffff != 0 {
				return 0, ErrNotUint32
			}
			break
		}
	}

	return uint32(output), nil
}
