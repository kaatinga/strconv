package faststrconv

// intError - Error type based on int value.
type intError int

// Error returns error description.
func (err intError) Error() string {
	return errorDescriptions[err]
}

// errorDescriptions contains descriptions for the intError errors.
var errorDescriptions = []string{
	ErrNotUint32: "the input string is not an uint32 number",
	ErrNotUint16: "the input string is not an uint16 number",
	ErrNotByte:   "the input string is not an uint8 number",
}

const (
	ErrNotUint32 intError = iota
	ErrNotUint16
	ErrNotByte
)
