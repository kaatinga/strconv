package faststrconv

import (
	"math/rand"
	"strconv"
	"testing"
)

func Test2String(t *testing.T) {
	var num int32
	for i := 0; i < 5000; i++ {
		result := strconv.Itoa(int(num))
		var gotUint32, gotUint16, gotByte string
		t.Run(result, func(t *testing.T) {
			gotUint32 = Uint322String(uint32(num))
			if gotUint32 != result {
				t.Errorf("Uint322String() = %v, want %v", gotUint32, result)
			}
			if num < 65535 {
				gotUint16 = Uint162String(uint16(num))
				if gotUint16 != result {
					t.Errorf("Uint162String() = %v, want %v", gotUint16, result)
				}
			}
			if num < 256 {
				gotByte = Byte2String(byte(num))
				if gotByte != result {
					t.Errorf("Byte2String() = %v, want %v", gotByte, result)
				}
			}
		})
		if num > 65535 {
			num = rand.Int31n(1535)
		} else {
			num = rand.Int31()
		}
	}
}
