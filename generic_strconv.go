package faststrconv

type stringOrBytes interface {
	~[]byte | ~string
}

// GetByte validates and converts input string or []byte to byte type.
func GetByte[I stringOrBytes](input I) (byte, error) {
	if len(input)&^byteLengthMask != 0 || len(input) == 0 {
		return 0, ErrNotByte
	}

	var i int
	var output uint16
	for {
		if input[i]&^digitsMask > 9 {
			return 0, ErrNotByte
		}

		output = uint16(input[i])&^digitsMask + output*10
		i++

		if i == len(input) {
			if output&^maxByteMask != 0 {
				return 0, ErrNotByte
			}
			break
		}
	}

	return byte(output), nil
}

// GetCustomByte validates and converts input string or []byte to byte type.
func GetCustomByte[I stringOrBytes, V ~byte](input I) (V, error) {
	if len(input)&^byteLengthMask != 0 || len(input) == 0 {
		return 0, ErrNotByte
	}

	var i int
	var output uint16
	for {
		if input[i]&^digitsMask > 9 {
			return 0, ErrNotByte
		}

		output = uint16(input[i])&^digitsMask + output*10
		i++

		if i == len(input) {
			if output&^maxByteMask != 0 {
				return 0, ErrNotByte
			}
			break
		}
	}

	return V(output), nil
}

// GetUint16 validates and converts input string or []byte to uint16 type.
func GetUint16[I stringOrBytes](input I) (uint16, error) {
	if len(input) == 0 {
		return 0, ErrNotUint16
	}

	var i int
	var output uint32
	for {
		if input[i]&^digitsMask > 9 {
			return 0, ErrNotUint16
		}

		output = output*10 + uint32(input[i])&^digitsMask
		i++

		if i == len(input) {
			if output&^maxUint16Mask != 0 {
				return 0, ErrNotUint16
			}
			break
		}
	}

	return uint16(output), nil
}

// GetUint32 validates and converts input string or []byte to uint32 type.
func GetUint32[I stringOrBytes](input I) (uint32, error) {
	if len(input) == 0 {
		return 0, ErrNotUint32
	}

	var i int
	var output uint64
	for {
		if input[i]&^digitsMask > 9 {
			return 0, ErrNotUint32
		}

		output = uint64(input[i])&^digitsMask + output*10
		i++

		if i == len(input) {
			if output&^maxUint32Mask != 0 {
				return 0, ErrNotUint32
			}
			break
		}
	}

	return uint32(output), nil
}
