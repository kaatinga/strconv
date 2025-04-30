package faststrconv

import (
	"strconv"
	"testing"
)

// //nolint
//	func BenchmarkGetByte(b *testing.B) {
//		b.ReportAllocs()
//		for i := 0; i < b.N; i++ {
//			GetByte("0")
//			//goland:noinspection GoUnhandledErrorResult
//			GetByte("255")
//		}
//	}
//
// // nolint
//	func BenchmarkGetUint16(b *testing.B) {
//		b.ReportAllocs()
//		for i := 0; i < b.N; i++ {
//			//goland:noinspection GoUnhandledErrorResult
//			GetUint16("0")
//			GetUint16("255")
//		}
//	}
//
// // nolint
//	func BenchmarkGetUint16DeletingLeadingZeros(b *testing.B) {
//		b.ReportAllocs()
//		for i := 0; i < b.N; i++ {
//			//goland:noinspection GoUnhandledErrorResult
//			GetUint16(DeleteLeadingZeros("0000"))
//			GetUint16(DeleteLeadingZeros("0000255"))
//		}
//	}
//
// // nolint
// func BenchmarkGetUint32(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		GetUint32("0")
// 		GetUint32("255")
// 		GetUint32("2550000")
// 		GetUint32("25500999900")
// 	}
// }
//
// // nolint
// func BenchmarkStrvconv_Atoi(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		//goland:noinspection ALL,GoUnhandledErrorResult
//		strconv.Atoi("0")
//		strconv.Atoi("255")
//	}
// }
//
// // nolint
// func BenchmarkStrvconv_AtoiDeletingLeadingZeros(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		//goland:noinspection ALL,GoUnhandledErrorResult
//		strconv.Atoi("0000")
//		strconv.Atoi("0000255")
//	}
// }

// func BenchmarkByte2String(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		_ = Byte2String(199)
// 		_ = Byte2String(0)
// 		_ = Byte2String(55)
// 	}
// }

// func BenchmarkUint162String(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		_ = Uint162String(199)
// 		_ = Uint162String(0)
// 		_ = Uint162String(55)
// 	}
// }

func BenchmarkByte2Bytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Byte2Bytes(199)
		_ = Byte2Bytes(0)
		_ = Byte2Bytes(55)
	}
}
func BenchmarkUint162Bytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Uint162Bytes(199)
		_ = Uint162Bytes(0)
		_ = Uint162Bytes(55)
		_ = Uint162Bytes(64555)
	}
}

// func BenchmarkUint2Bytes(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		_ = Uint2String(199)
// 		_ = Uint2String(0)
// 		_ = Uint2String(55)
// 		_ = Uint2String(64555)
// 	}
// }

//	func BenchmarkUint322String(b *testing.B) {
//		b.ReportAllocs()
//		for i := 0; i < b.N; i++ {
//			Uint322String(199)
//			Uint322String(0)
//			Uint322String(55)
//			Uint322String(64555)
//		}
//	}
//
//	func BenchmarkUint642String(b *testing.B) {
//		b.ReportAllocs()
//		for i := 0; i < b.N; i++ {
//			Uint642String(199)
//			Uint642String(0)
//			Uint642String(55)
//			Uint642String(64555)
//		}
//	}
//
//	func BenchmarkUint32And642StringWithUint64(b *testing.B) {
//		uint199 := uint64(199)
//		uint0 := uint64(0)
//		uint55 := uint64(55)
//		uint64555 := uint64(64555)
//		b.ReportAllocs()
//		for i := 0; i < b.N; i++ {
//			Uint32And642String(uint199)
//			Uint32And642String(uint0)
//			Uint32And642String(uint55)
//			Uint32And642String(uint64555)
//		}
//	}
//
//	func BenchmarkUint32And642StringWithUint32(b *testing.B) {
//		b.ReportAllocs()
//		for i := 0; i < b.N; i++ {
//			Uint32And642String(uint32(199))
//			Uint32And642String(uint32(0))
//			Uint32And642String(uint32(55))
//			Uint32And642String(uint32(64555))
//		}
//	}
func BenchmarkItoa(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(199)
		strconv.Itoa(0)
		strconv.Itoa(55)
		strconv.Itoa(64555)
	}
}
