package strconv

const (
	UnicodeMask        = 0xf
	ByteLengthMask int = 0b11 // it allows to catch cases faster when the number is too long.
)
