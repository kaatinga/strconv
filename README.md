[![GitHub release](https://img.shields.io/github/release/kaatinga/strconv.svg)](https://github.com/kaatinga/strconv/releases)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kaatinga/strconv/blob/main/LICENSE)
[![codecov](https://codecov.io/gh/kaatinga/strconv/branch/master/graph/badge.svg)](https://codecov.io/gh/kaatinga/strconv)
[![lint workflow](https://github.com/kaatinga/strconv/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/kaatinga/strconv/actions?query=workflow%3Alinter)
[![help wanted](https://img.shields.io/badge/Help%20wanted-True-yellow.svg)](https://github.com/kaatinga/strconv/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)

# The strconv package

The package contain several useful functions intended for conversion between string,
[]byte and different integer types.

For the sake of code code shortness, generics are used, it means the package requires to use go1.18+.

## GetUint32(), GetUint16(),GetByte()

The functions are faster alternative to `strconv.Atoi()`.

```
goos: darwin
goarch: arm64
pkg: github.com/kaatinga/strconv
BenchmarkGetByte
BenchmarkGetByte-8         	420296895	         2.592 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetUint16
BenchmarkGetUint16-8       	477782215	         2.521 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetUint32
BenchmarkGetUint32-8       	473668162	         2.518 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrvconv_Atoi
BenchmarkStrvconv_Atoi-8   	180559644	         6.732 ns/op	       0 B/op	       0 allocs/op
PASS
```

## Uint162String, Byte2String, etc

The functions are faster alternative to `strconv.Itoa()`.

```
cpu: AMD Ryzen 5 3400G with Radeon Vega Graphics    
BenchmarkUint162String
BenchmarkUint162String-8   	68361266	        17.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrvconvItoa
BenchmarkStrvconvItoa-8    	30279402	        39.60 ns/op	       5 B/op	       1 allocs/op
```

Warning. For the sake of maximum processing speed, all the converters have limited support of the leading zeros.
The length of the string cannot exceed the maximum length of the string of every certain type. For example, `012` is
supported for the `byte` type, whereas `0255` is not as it consists of 4 characters what is impossible for `byte` type.
In case you process a date number, like month that is formatted with zero, like `09`, you still can use the fast
converters of the strconv package directly, but if you still need to process formatted numbers like `0000000000000255` as byte, use
`DeleteLeadingZeros` helper to remove leading zeros in the processed string first:

```go
numberINeed, err = strconv.String2Byte(DeleteLeadingZeros("0000000000000255"))
if err != nil {
	...
}
```