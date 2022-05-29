package faststrconv

// Byte2String converts a byte number to string.
func Byte2String(num byte) string {
	return string(Byte2Bytes(num))
}

// Uint162String converts an uint16 number to string.
func Uint162String(num uint16) string {
	return string(Uint162Bytes(num))
}

// Uint322String converts an uint32 number to string.
func Uint322String(num uint32) string {
	return string(Uint322Bytes(num))
}
