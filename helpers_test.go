package strconv

import (
	"reflect"
	"testing"
)

func TestDeleteLeadingZeros(t *testing.T) {
	tests := []struct {
		numberAsString []byte
		want           []byte
	}{
		{[]byte("000022"), []byte("22")},
		{[]byte("0000"), []byte("0")},
	}
	for _, tt := range tests {
		t.Run(string(tt.numberAsString), func(t *testing.T) {
			if got := DeleteLeadingZeros(tt.numberAsString); !reflect.DeepEqual(got, tt.want) { // nolint:scopelint
				t.Errorf("DeleteLeadingZeros() = %v, want %v", got, tt.want) // nolint:scopelint
			}
		})
	}
}
