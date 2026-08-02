package faststrconv

import "testing"

func TestIntErrorOutOfRange(t *testing.T) {
	const want = "unknown integer conversion error"
	for _, err := range []intError{ErrNotUint32 - 1, ErrNotByte + 1} {
		if got := err.Error(); got != want {
			t.Errorf("intError(%d).Error() = %q, want %q", err, got, want)
		}
	}
}
