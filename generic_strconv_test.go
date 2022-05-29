package faststrconv

import (
	"strconv"
	"testing"
)

func Test2Byte(t *testing.T) {
	var stringValue string
	var value int
	var value3 byte
	var err error
	for i := 0; i < 1000; i++ {
		stringValue = strconv.Itoa(i)
		value, _ = strconv.Atoi(stringValue)

		value3, err = GetByte(stringValue)
		if i < 256 {
			if err != nil {
				t.Errorf("String2Byte() error = %v, want %v", err, nil)
			}

			if byte(value) != value3 {
				t.Errorf("String2Byte() value = %v, want %v", value3, value)
			}
		} else {
			if err == nil {
				t.Error("String2Byte() must have an error")
			}

			if 0 != value3 {
				t.Errorf("String2Byte() value = %v, want %v, i %v", value3, 0, i)
			}
		}
	}
}

func TestGetUint32(t *testing.T) {
	var stringValue string
	var value int
	var value3 uint32
	var err error
	for i := uint64(0); i < 10000000000; i = i + 10001 {
		stringValue = strconv.FormatUint(i, 10)
		value, _ = strconv.Atoi(stringValue)

		value3, err = GetUint32(stringValue)
		if i < MaxUint32 {
			if err != nil {
				t.Errorf("String2Uint32() error = %v, want %v", err, nil)
			}

			if uint32(value) != value3 {
				t.Errorf("String2Uint32() value = %v, want %v", value3, value)
			}
		} else {
			if err == nil {
				t.Error("String2Uint32() must have an error")
			}

			if 0 != value3 {
				t.Errorf("String2Uint32() value = %v, want 0, i = %v", value3, i)
			}
		}
	}
}

func TestGetByteAndGetUint16(t *testing.T) {
	tests := []struct {
		input   []byte
		want    uint16
		wantErr bool
	}{
		{[]byte("255"), 255, false},
		{[]byte("25"), 25, false},
		{[]byte("025"), 25, false},
		{[]byte("0"), 0, false},
		{[]byte("16000"), 16000, false},
		{[]byte("66535"), 0, true},
		{[]byte("160000"), 0, true},
		{[]byte("abc"), 0, true},
		{[]byte(""), 0, true},
	}
	for _, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {
			gotUint16, err := GetUint16(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUint16 != tt.want {
				t.Errorf("GetUint16() gotUint16 = %v, want %v", gotUint16, tt.want)
			}
			if string(tt.input) != "16000" {
				gotByte, err := GetByte(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetByte() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotByte != byte(tt.want) {
					t.Errorf("GetByte() gotUint16 = %v, want %v", gotByte, tt.want)
				}
			}
			if string(tt.input) == "" || string(tt.input) == "abc" {
				gotUint32, err := GetUint32(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("gotUint32 error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotUint32 != uint32(tt.want) {
					t.Errorf("gotUint32 gotUint16 = %v, want %v", gotUint32, tt.want)
				}
			}
		})
	}
}
