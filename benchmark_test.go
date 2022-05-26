package faststrconv

import (
	"strconv"
	"testing"
)

//nolint
func BenchmarkGetByte(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetByte("0")
		//goland:noinspection GoUnhandledErrorResult
		GetByte("255")
	}
}

// nolint
func BenchmarkGetUint16(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		//goland:noinspection GoUnhandledErrorResult
		GetUint16("0")
		GetUint16("255")
	}
}

// nolint
func BenchmarkGetUint16DeletingLeadingZeros(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		//goland:noinspection GoUnhandledErrorResult
		GetUint16(DeleteLeadingZeros("0000"))
		GetUint16(DeleteLeadingZeros("0000255"))
	}
}

// nolint
func BenchmarkGetUint32(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetUint32("0")
		GetUint32("255")
	}
}

// nolint
func BenchmarkStrvconv_Atoi(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		//goland:noinspection ALL,GoUnhandledErrorResult
		strconv.Atoi("0")
		strconv.Atoi("255")
	}
}

// nolint
func BenchmarkStrvconv_AtoiDeletingLeadingZeros(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		//goland:noinspection ALL,GoUnhandledErrorResult
		strconv.Atoi("0000")
		strconv.Atoi("0000255")
	}
}
