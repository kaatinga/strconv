package faststrconv

const (
	MaxByteMask    = 0xff
	MaxUint16Mask  = 0xffff
	MaxUint32Mask  = 0xffffffff
	DigitsMask     = 0b110000
	ByteLengthMask = 0b11 // it allows to catch cases faster when the number is too long.
)
