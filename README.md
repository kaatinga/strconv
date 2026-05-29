[![CodeQL](https://github.com/kaatinga/strconv/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/kaatinga/strconv/actions/workflows/github-code-scanning/codeql)
[![GitHub release](https://img.shields.io/github/release/kaatinga/strconv.svg)](https://github.com/kaatinga/strconv/releases)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kaatinga/strconv/blob/main/LICENSE)
[![codecov](https://codecov.io/gh/kaatinga/strconv/branch/main/graph/badge.svg?token=TL88FINYP4)](https://codecov.io/gh/kaatinga/strconv)
[![Go Tests](https://github.com/kaatinga/strconv/actions/workflows/golang_ci.yml/badge.svg)](https://github.com/kaatinga/strconv/actions/workflows/golang_ci.yml)
[![help wanted](https://img.shields.io/badge/Help%20wanted-True-yellow.svg)](https://github.com/kaatinga/strconv/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)

# faststrconv

A small, dependency-free package with allocation-light helpers for converting
between `string`, `[]byte` and unsigned integer types (`byte`, `uint16`,
`uint32`, `uint64`).

It is built for the hot path: parsing is **~2–2.5× faster than the standard
library with zero allocations**, and the integer → `[]byte` converters are
faster than `strconv.AppendUint` while allocating an exactly-sized slice.

Generics are used for the type-parametrised helpers, so the package requires
**Go 1.18+**.

```go
import faststrconv "github.com/kaatinga/strconv"
```

## Benchmarks

All numbers below are `ns/op` (lower is better) and `B/op` from
`go test -bench=. -benchmem`, best of 8 runs on:

```
goos: darwin   goarch: arm64   cpu: Apple M1   go: 1.26
```

Methodology note: inputs are read from mutable variables and results are stored
in package-level sinks, otherwise the compiler constant-folds the calls and
escape analysis removes the allocations — which is why naive micro-benchmarks of
this kind tend to report a misleading `0 ns/op, 0 allocs/op`. The benchmarks
live in [`benchmark_test.go`](./benchmark_test.go) and are paired with the
closest `strconv` equivalent over identical inputs.

### Parsing: `string` / `[]byte` → integer

`GetByte`, `GetUint16` and `GetUint32` are a faster, allocation-free
alternative to `strconv.ParseUint` / `strconv.Atoi`.

| Function     | faststrconv         | `strconv.ParseUint` | Speed-up |
|--------------|---------------------|---------------------|----------|
| → `byte`     | `GetByte`   **2.6 ns**, 0 B | 6.0 ns, 0 B | **~2.3×** |
| → `uint16`   | `GetUint16` **3.2 ns**, 0 B | 8.1 ns, 0 B | **~2.5×** |
| → `uint32`   | `GetUint32` **4.8 ns**, 0 B | 10.1 ns, 0 B | **~2.1×** |

All three accept both `string` and `[]byte` input (generic over
`~string | ~[]byte`) and never allocate.

### Formatting: integer → `[]byte`

The `*2Bytes` helpers return a fresh, exactly-sized slice. The closest standard
library equivalent is `strconv.AppendUint(nil, x, 10)`.

| Function       | faststrconv          | `strconv.AppendUint(nil, …)` |
|----------------|----------------------|------------------------------|
| `Byte2Bytes`   | **10.0 ns**, 2 B     | 26 ns, 12 B |
| `Uint162Bytes` | **12.4 ns**, 4 B     | 26 ns, 12 B |
| `Uint322Bytes` | **15.0 ns**, 8 B     | 26 ns, 12 B |
| `Uint642Bytes` | **19.1 ns**, 10 B    | 26 ns, 12 B |
| `Uint2Bytes`   | **19.3 ns**, 10 B    | 26 ns, 12 B |

Every converter is both faster and lighter on memory, and the returned slice is
trimmed to the exact digit count (no oversized backing array).

### Formatting: integer → `string`

| Function        | faststrconv       | `strconv` (`FormatUint`/`Itoa`) |
|-----------------|-------------------|---------------------------------|
| `Byte2String`   | **9.5 ns**, 2 B   | 15–17 ns |
| `Uint162String` | **12.2 ns**, 2 B  | 15–17 ns |
| `Uint322String` | 24.8 ns, 12 B     | ~17 ns |
| `Uint642String` | 29.5 ns, 18 B     | ~17 ns |

For `byte` and `uint16` the string converters comfortably beat the standard
library. For `uint32`/`uint64`, `strconv.FormatUint` is competitive because it
keeps an internal cache of small-integer strings; if you need the fastest path
for large integers, prefer the `*2Bytes` converters above (or `strconv`
directly).

## API

### Parsing

```go
b,   err := faststrconv.GetByte("255")    // byte
u16, err := faststrconv.GetUint16("64555") // uint16
u32, err := faststrconv.GetUint32("4294967295") // uint32

// Works with []byte too, and with custom types whose underlying type is byte:
v, err := faststrconv.GetCustomByte[string, MyByte]("42")
```

### Formatting

```go
faststrconv.Byte2Bytes(255)        // []byte("255")
faststrconv.Uint162Bytes(64555)    // []byte
faststrconv.Uint322Bytes(123456)   // []byte
faststrconv.Uint642Bytes(123456789)// []byte
faststrconv.Uint2Bytes(uint64(42)) // generic over uint | uint32 | uint64

faststrconv.Byte2String(255)        // "255"
faststrconv.Uint162String(64555)    // "64555"
faststrconv.Uint322String(123456)   // "123456"
faststrconv.Uint642String(123456789)// "123456789"
faststrconv.Uint32And642String(uint64(42)) // generic
```

## Leading zeros

> [!WARNING]
> For the sake of maximum processing speed, the parsers have limited support of
> leading zeros: the input length cannot exceed the maximum digit length of the
> target type. For example `012` is accepted for `byte`, but `0255` is not — it
> is 4 characters long, which is impossible for a `byte` value.

If you need to parse a zero-padded number such as `0000000000000255`, strip the
leading zeros first with the `DeleteLeadingZeros` helper:

```go
numberINeed, err := faststrconv.GetByte(faststrconv.DeleteLeadingZeros("0000000000000255"))
if err != nil {
	// ...
}
```
