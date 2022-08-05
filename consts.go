package faststrconv

const (
	maxByteMask    = 0xff
	maxUint16Mask  = 0xffff
	maxUint32Mask  = 0xffffffff
	digitsMask     = 0b110000
	byteLengthMask = 0b11 // it allows to catch cases faster when the number is too long.

	maxUint8  = 1<<8 - 1
	maxUint16 = 1<<16 - 1
	maxUint32 = 1<<32 - 1
)
