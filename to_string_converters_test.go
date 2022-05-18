package strconv

import (
	"reflect"
	"testing"
)

func TestUint162String(t *testing.T) {

	tests := []struct {
		result string
		num    uint16
	}{
		{"199", 199},
		{"1999", 1999},
		{"222", 222},
		{"1", 1},
		{"0", 0},
		{"55555", 55555},
		{"12345", 12345},
		{"10000", 10000},
		{"9999", 9999},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.result, func(t *testing.T) {
			got := Uint162String(tt.num)
			if got != tt.result {
				t.Errorf("Uint162String() = %v, want %v", got, tt.result)
			}
		})
	}
}

func TestByte2String(t *testing.T) {

	tests := []struct {
		result string
		num    byte
	}{
		{"199", 199},
		{"99", 99},
		{"100", 100},
		{"255", 255},
		{"222", 222},
		{"0", 0},
		{"9", 9},
		{"10", 10},
		{"1", 1},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.result, func(t *testing.T) {
			if got := Byte2String(tt.num); !reflect.DeepEqual(got, tt.result) {
				t.Errorf("Byte2String() = %v, want %v", got, tt.result)
			}
		})
	}
}
