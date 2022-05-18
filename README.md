# The strconv package

The package contain several useful functions intended for conversion between string,
[]byte and different integer types.

## String2Uint32(), Bytes2Uint32(), String2Byte(), etc

The functions are faster alternative to `strconv.Atoi()`.

```
goos: darwin
goarch: arm64
pkg: github.com/kaatinga/strconv
BenchmarkString2Byte
BenchmarkString2Byte-8      	467720466	         2.542 ns/op	       0 B/op	       0 allocs/op
BenchmarkGeneric2Uint16
BenchmarkGeneric2Uint16-8   	476605455	         2.535 ns/op	       0 B/op	       0 allocs/op
BenchmarkString2Uint32
BenchmarkString2Uint32-8    	473442506	         2.533 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrvconv_Atoi
BenchmarkStrvconv_Atoi-8    	176394312	         6.808 ns/op	       0 B/op	       0 allocs/op
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
converters of the assets package directly, but if you still need to process formatted numbers like `0000000000000255` as byte, use
`DeleteLeadingZeros` helper to remove leading zeros in the processed string first:

```go
numberINeed, err = assets.String2Byte(DeleteLeadingZeros("0000000000000255"))
if err != nil {
	...
}
```