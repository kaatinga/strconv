[![Tests](https://github.com/kaatinga/strconv/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/kaatinga/strconv/actions/workflows/test.yml)
[![GitHub release](https://img.shields.io/github/release/kaatinga/strconv.svg)](https://github.com/kaatinga/strconv/releases)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kaatinga/strconv/blob/main/LICENSE)
[![codecov](https://codecov.io/gh/kaatinga/strconv/branch/main/graph/badge.svg?token=TL88FINYP4)](https://codecov.io/gh/kaatinga/strconv)
[![lint workflow](https://github.com/kaatinga/strconv/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/kaatinga/strconv/actions?query=workflow%3Alinter)
[![help wanted](https://img.shields.io/badge/Help%20wanted-True-yellow.svg)](https://github.com/kaatinga/strconv/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)

# The faststrconv package

The package contain several useful functions intended for conversion between string,
[]byte and different integer types.

For the sake of code shortness, generics are used, it means the package requires go1.18+.

## GetUint32(), GetUint16(), GetByte()

The functions are faster alternative to `strconv.Atoi()`.

```
goos: darwin
goarch: arm64
pkg: github.com/kaatinga/strconv
BenchmarkGetByte
BenchmarkGetByte-8                             	523370584	         2.042 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetUint16
BenchmarkGetUint16-8                           	582942614	         2.046 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetUint32
BenchmarkGetUint32-8                           	587540109	         2.045 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrvconv_Atoi
BenchmarkStrvconv_Atoi-8                       	174534008	         7.387 ns/op	       0 B/op	       0 allocs/op
PASS
```

## Uint162String, Byte2String, etc

The functions are faster alternative to `strconv.Itoa()`.

```
goos: darwin
goarch: arm64
pkg: github.com/kaatinga/strconv
BenchmarkByte2String
BenchmarkByte2String-8                    	100000000	        10.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint162String
BenchmarkUint162String-8                  	71960958	        16.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint322String
BenchmarkUint322String-8                  	71680485	        16.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint642String
BenchmarkUint642String-8                  	65918067	        17.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint32And642StringWithUint64
BenchmarkUint32And642StringWithUint64-8   	58261146	        19.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint32And642StringWithUint32
BenchmarkUint32And642StringWithUint32-8   	65289982	        18.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkItoa
BenchmarkItoa-8                           	34061032	        35.20 ns/op	       8 B/op	       2 allocs/op
PASS
```

Warning. For the sake of maximum processing speed, all the converters have limited support of the leading zeros.
The length of the string cannot exceed the maximum length of the string of every certain type. For example, `012` is
supported for the `byte` type, whereas `0255` is not as it consists of 4 characters what is impossible for `byte` type.
In case you process a date number, like month that is formatted with zero, like `09`, you can use the fast
converters of the strconv package directly, but if you still need to process formatted numbers like `0000000000000255` as byte, use
`DeleteLeadingZeros` helper to remove leading zeros in the processed string first:

```go
numberINeed, err = strconv.String2Byte(DeleteLeadingZeros("0000000000000255"))
if err != nil {
	...
}
```