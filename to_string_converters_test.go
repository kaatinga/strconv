package faststrconv

import (
	"math/rand"
	"strconv"
	"testing"
)

func Test2String(t *testing.T) {
	var num uint64
	for i := 0; i < 70; i++ {
		result := strconv.FormatUint(num, 10)
		var got string
		t.Run(result, func(t *testing.T) {
			got = Uint642String(num)
			if got != result {
				t.Errorf("Uint642String() = %v, want %v", got, result)
			}
			got = Uint32And642String(num)
			if got != result {
				t.Errorf("Uint2Bytes() = %v, want %v", got, result)
			}
			if num < MaxUint32 {
				got = Uint322String(uint32(num))
				if got != result {
					t.Errorf("Uint322String() = %v, want %v", got, result)
				}
			}
			if num < MaxUint16 {
				got = Uint162String(uint16(num))
				if got != result {
					t.Errorf("Uint162String() = %v, want %v", got, result)
				}
			}
			if num < MaxUint8 {
				got = Byte2String(byte(num))
				if got != result {
					t.Errorf("Byte2String() = %v, want %v", got, result)
				}
			}
		})
		num += uint64(rand.Int31n(10)) + num
	}
}
