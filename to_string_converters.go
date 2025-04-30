package faststrconv

// Byte2String converts a byte to string.
func Byte2String(num byte) string {
	return string(Byte2Bytes(num))
}

// Uint162String converts an uint16 to string.
func Uint162String(num uint16) string {
	return string(Uint162Bytes(num))
}

// Uint322String converts an uint32 to string.
func Uint322String(num uint32) string {
	return string(Uint322Bytes(num))
}

// Uint642String converts an uint64 to string.
func Uint642String(num uint64) string {
	return string(Uint642Bytes(num))
}

// Uint32And642String converts an uint64 or uint32 to string.
func Uint32And642String[UI unsigned](num UI) string {
	return string(Uint2Bytes(num))
}
