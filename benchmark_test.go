package faststrconv

import (
	"strconv"
	"testing"
)

////nolint
//func BenchmarkGetByte(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		GetByte("0")
//		//goland:noinspection GoUnhandledErrorResult
//		GetByte("255")
//	}
//}
//
//// nolint
//func BenchmarkGetUint16(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		//goland:noinspection GoUnhandledErrorResult
//		GetUint16("0")
//		GetUint16("255")
//	}
//}
//
//// nolint
//func BenchmarkGetUint16DeletingLeadingZeros(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		//goland:noinspection GoUnhandledErrorResult
//		GetUint16(DeleteLeadingZeros("0000"))
//		GetUint16(DeleteLeadingZeros("0000255"))
//	}
//}
//
//// nolint
//func BenchmarkGetUint32(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		GetUint32("0")
//		GetUint32("255")
//	}
//}
//
//// nolint
//func BenchmarkStrvconv_Atoi(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		//goland:noinspection ALL,GoUnhandledErrorResult
//		strconv.Atoi("0")
//		strconv.Atoi("255")
//	}
//}
//
//// nolint
//func BenchmarkStrvconv_AtoiDeletingLeadingZeros(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		//goland:noinspection ALL,GoUnhandledErrorResult
//		strconv.Atoi("0000")
//		strconv.Atoi("0000255")
//	}
//}

func BenchmarkByte2String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Byte2String(199)
		Byte2String(0)
		Byte2String(55)
	}
}

func BenchmarkUint162String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint162String(199)
		Uint162String(0)
		Uint162String(55)
		Uint162String(64555)
	}
}

func BenchmarkUint322String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint322String(199)
		Uint322String(0)
		Uint322String(55)
		Uint322String(55)
		Uint322String(64555)
	}
}

func BenchmarkItoa(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(199)
		strconv.Itoa(0)
		strconv.Itoa(55)
		strconv.Itoa(55)
		strconv.Itoa(64555)
	}
}
