package strconv

// IntError - Error type based on int value.
type IntError int

// Error returns error description.
func (err IntError) Error() string {
	return errorDescriptions[err]
}

// errorDescriptions contains descriptions for the IntError errors.
var errorDescriptions = []string{
	ErrNotUint32: "the input string is not an uint32 number",
	ErrNotUint16: "the input string is not an uint16 number",
	ErrNotByte:   "the input string is not an uint8 number",
}

const (
	ErrNotUint32 IntError = iota
	ErrNotUint16
	ErrNotByte
)
