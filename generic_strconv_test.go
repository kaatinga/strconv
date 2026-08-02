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

		t.Run("GetByte + "+stringValue, func(t *testing.T) {
			value3, err = GetByte(stringValue)
			if i < 256 { //nolint:nestif
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

				if value3 != 0 {
					t.Errorf("String2Byte() value = %v, want %v, i %v", value3, 0, i)
				}
			}
		})

		t.Run("GetCustomByte + "+stringValue, func(t *testing.T) {
			value3, err = GetCustomByte[string, byte](stringValue)
			if i < 256 { //nolint:nestif
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

				if value3 != 0 {
					t.Errorf("String2Byte() value = %v, want %v, i %v", value3, 0, i)
				}
			}
		})
	}
}

func TestGetUint32(t *testing.T) {
	var stringValue string
	var value3 uint32
	var err error
	for i := uint64(0); i < 10000000000; i = i + 10001 {
		stringValue = strconv.FormatUint(i, 10)

		value3, err = GetUint32(stringValue)
		if i <= maxUint32 { //nolint:nestif
			if err != nil {
				t.Errorf("String2Uint32() error = %v, want %v", err, nil)
			}

			if uint32(i) != value3 {
				t.Errorf("String2Uint32() value = %v, want %v", value3, i)
			}
		} else {
			if err == nil {
				t.Error("String2Uint32() must have an error")
			}

			if value3 != 0 {
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
	for index := range tests {
		t.Run(string(tests[index].input), func(t *testing.T) {
			gotUint16, err := GetUint16(tests[index].input)
			if (err != nil) != tests[index].wantErr {
				t.Errorf("GetUint16() error = %v, wantErr %v", err, tests[index].wantErr)
				return
			}
			if gotUint16 != tests[index].want {
				t.Errorf("GetUint16() gotUint16 = %v, want %v", gotUint16, tests[index].want)
			}
			if string(tests[index].input) != "16000" {
				gotByte, err := GetByte(tests[index].input)
				if (err != nil) != tests[index].wantErr {
					t.Errorf("GetByte() error = %v, wantErr %v", err, tests[index].wantErr)
					return
				}
				if gotByte != byte(tests[index].want) {
					t.Errorf("GetByte() gotUint16 = %v, want %v", gotByte, tests[index].want)
				}
			}
			if string(tests[index].input) == "" || string(tests[index].input) == "abc" {
				gotUint32, err := GetUint32(tests[index].input)
				if (err != nil) != tests[index].wantErr {
					t.Errorf("gotUint32 error = %v, wantErr %v", err, tests[index].wantErr)
					return
				}
				if gotUint32 != uint32(tests[index].want) {
					t.Errorf("gotUint32 gotUint16 = %v, want %v", gotUint32, tests[index].want)
				}
			}
		})
	}
}

func TestParsersRejectInvalidDigits(t *testing.T) {
	inputs := []string{"1 "}
	for value := 0; value <= 255; value++ {
		if value < '0' || value > '9' {
			inputs = append(inputs, string([]byte{byte(value)}))
		}
	}

	for _, input := range inputs {
		t.Run(strconv.Quote(input), func(t *testing.T) {
			if _, err := GetByte(input); err != ErrNotByte {
				t.Errorf("GetByte() error = %v, want %v", err, ErrNotByte)
			}
			if _, err := GetCustomByte[string, byte](input); err != ErrNotByte {
				t.Errorf("GetCustomByte() error = %v, want %v", err, ErrNotByte)
			}
			if _, err := GetUint16(input); err != ErrNotUint16 {
				t.Errorf("GetUint16() error = %v, want %v", err, ErrNotUint16)
			}
			if _, err := GetUint32(input); err != ErrNotUint32 {
				t.Errorf("GetUint32() error = %v, want %v", err, ErrNotUint32)
			}
		})
	}
}

func TestUintParsersBounds(t *testing.T) {
	t.Run("uint16", func(t *testing.T) {
		tests := []struct {
			input   string
			want    uint16
			wantErr error
		}{
			{"65535", maxUint16, nil},
			{"65536", 0, ErrNotUint16},
			{"000001", 0, ErrNotUint16},
			{"4294967296", 0, ErrNotUint16},
		}
		for _, tt := range tests {
			got, err := GetUint16(tt.input)
			if err != tt.wantErr {
				t.Errorf("GetUint16(%q) error = %v, want %v", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("GetUint16(%q) = %v, want %v", tt.input, got, tt.want)
			}
		}
	})

	t.Run("uint32", func(t *testing.T) {
		tests := []struct {
			input   string
			want    uint32
			wantErr error
		}{
			{"4294967295", maxUint32, nil},
			{"4294967296", 0, ErrNotUint32},
			{"00000000001", 0, ErrNotUint32},
			{"18446744073709551616", 0, ErrNotUint32},
		}
		for _, tt := range tests {
			got, err := GetUint32(tt.input)
			if err != tt.wantErr {
				t.Errorf("GetUint32(%q) error = %v, want %v", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("GetUint32(%q) = %v, want %v", tt.input, got, tt.want)
			}
		}
	})
}
