package faststrconv

import (
	"math/rand"
	"strconv"
	"testing"
)

func Test2String(t *testing.T) {
	var num int32
	for i := 0; i < 1000; i++ {
		result := strconv.Itoa(int(num))
		t.Run(result, func(t *testing.T) {
			gotUint16 := Uint162String(uint16(num))
			if gotUint16 != result {
				t.Errorf("Uint162String() = %v, want %v", gotUint16, result)
			}
			if num < 256 {
				gotByte := Byte2String(byte(num))
				if gotByte != result {
					t.Errorf("Byte2String() = %v, want %v", gotByte, result)
				}
			}
		})
		num = rand.Int31n(65535)
	}
}
