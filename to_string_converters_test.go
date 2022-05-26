package faststrconv

import (
	"testing"
)

func Test2String(t *testing.T) {
	tests := []struct {
		result string
		num    uint16
	}{
		{"0", 0},
		{"1", 1},
		{"9", 9},
		{"10", 10},
		{"22", 22},
		{"99", 99},
		{"100", 100},
		{"199", 199},
		{"222", 222},
		{"255", 255},
		{"1999", 1999},
		{"9999", 9999},
		{"10000", 10000},
		{"55555", 55555},
		{"12345", 12345},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.result, func(t *testing.T) {
			gotUint16 := Uint162String(tt.num)
			if gotUint16 != tt.result {
				t.Errorf("Uint162String() = %v, want %v", gotUint16, tt.result)
			}
			if tt.num < 256 {
				gotByte := Byte2String(byte(tt.num))
				if gotByte != tt.result {
					t.Errorf("Byte2String() = %v, want %v", gotByte, tt.result)
				}
			}
		})
	}
}
