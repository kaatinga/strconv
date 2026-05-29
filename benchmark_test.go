package faststrconv

import (
	"strconv"
	"testing"
)

// NOTE on methodology: benchmarks must read their inputs from mutable
// package-level slices (indexed by the loop counter) and store their results in
// package-level sinks. Passing compile-time constants lets the compiler
// constant-fold/elide the work, and discarding results lets escape analysis
// remove the allocations entirely - both make the numbers meaninglessly report
// "0 ns/op, 0 B/op, 0 allocs/op".
//
// Every faststrconv benchmark is paired with the closest standard-library
// strconv equivalent over the same inputs so the two can be compared directly.
var (
	sinkString string
	sinkBytes  []byte
	sinkU16    uint16
	sinkByte   byte
	sinkU32    uint32

	inBytes   = []byte{199, 0, 55, 255}
	inU16     = []uint16{199, 0, 55, 64555}
	inU32     = []uint32{199, 0, 2550000, 4294967295}
	inU64     = []uint64{199, 0, 2550000, 18446744073709551615}
	inParseB  = []string{"0", "99", "255", "7"}
	inParse16 = []string{"0", "255", "16000", "64555"}
	inParse32 = []string{"0", "255", "2550000", "4294967295"}
)

// ---------------------------------------------------------------------------
// string -> number (parsing)   vs   strconv.ParseUint
// ---------------------------------------------------------------------------

func BenchmarkGetByte(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkByte, _ = GetByte(inParseB[i&3])
	}
}

func BenchmarkParseUint_byte(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v, _ := strconv.ParseUint(inParseB[i&3], 10, 8)
		sinkByte = byte(v)
	}
}

func BenchmarkGetUint16(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkU16, _ = GetUint16(inParse16[i&3])
	}
}

func BenchmarkParseUint_uint16(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v, _ := strconv.ParseUint(inParse16[i&3], 10, 16)
		sinkU16 = uint16(v)
	}
}

func BenchmarkGetUint32(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkU32, _ = GetUint32(inParse32[i&3])
	}
}

func BenchmarkParseUint_uint32(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v, _ := strconv.ParseUint(inParse32[i&3], 10, 32)
		sinkU32 = uint32(v)
	}
}

func BenchmarkAtoi(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v, _ := strconv.Atoi(inParse32[i&3])
		sinkU32 = uint32(v)
	}
}

// ---------------------------------------------------------------------------
// number -> string   vs   strconv.FormatUint / strconv.Itoa
// ---------------------------------------------------------------------------

func BenchmarkByte2String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkString = Byte2String(inBytes[i&3])
	}
}

func BenchmarkUint162String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkString = Uint162String(inU16[i&3])
	}
}

func BenchmarkUint322String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkString = Uint322String(inU32[i&3])
	}
}

func BenchmarkUint642String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkString = Uint642String(inU64[i&3])
	}
}

func BenchmarkUint32And642String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkString = Uint32And642String(inU64[i&3])
	}
}

func BenchmarkFormatUint(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkString = strconv.FormatUint(inU64[i&3], 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkString = strconv.Itoa(int(inU32[i&3]))
	}
}

// ---------------------------------------------------------------------------
// number -> []byte   vs   strconv.AppendUint
// ---------------------------------------------------------------------------

func BenchmarkByte2Bytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkBytes = Byte2Bytes(inBytes[i&3])
	}
}

func BenchmarkUint162Bytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkBytes = Uint162Bytes(inU16[i&3])
	}
}

func BenchmarkUint322Bytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkBytes = Uint322Bytes(inU32[i&3])
	}
}

func BenchmarkUint642Bytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkBytes = Uint642Bytes(inU64[i&3])
	}
}

func BenchmarkUint2Bytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkBytes = Uint2Bytes(inU64[i&3])
	}
}

// AppendUint into a fresh, exactly-sized slice is the closest strconv analogue
// to the Uint*2Bytes helpers (which also return a fresh, exactly-sized slice).
func BenchmarkAppendUint(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkBytes = strconv.AppendUint(nil, inU64[i&3], 10)
	}
}
