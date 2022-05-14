## Basic Types

You can visit [runoob](https://www.runoob.com/go/go-data-types.html) to see more detailed definition.

### Integers
```go
// all numeric types default to 0

// unsigned int with 8 bits
// Can store: 0 to 255
var myint uint8
// signed int with 8 bits
// Can store: -127 to 127
var myint int8

// unsigned int with 16 bits
// Can store: 0 to 65535
var myint uint16
// signed int with 16 bits
// Can store: -32767 to 32767
var myint int16

// unsigned int with 32 bits
// Can store: 0 to 4294967295
var myint uint32
// signed int with 32 bits
// Can store: -2147483647 to 2147483647
var myint int32

// unsigned int with 64 bits
// Can store: 0 to 18446744073709551615
var myint uint64
// signed int with 64 bits
// Can store: -9223372036854775808 to 9223372036854775807
var myint int64
```

### Floating Point Numbers

These come in 2 distinct sizes, either `float32` or `float64`.

```go
var f1 float32
var f2 float64
```

Either `float32` or `float64` will allow you to work with exceptionally large numbers that don't fit inside a standard `int64` data type.

```go
var maxFloat32 float32
maxFloat32 = 16777216
fmt.Println(maxFloat32 == maxFloat32+10) // you would typically expect this to return false
// it returns true
fmt.Println(maxFloat32+10) // 16777216
fmt.Println(maxFloat32+2000000) // 16777216
```

### Complex Numbers

These come in 2 distinct sizes, either `complex64` or `complex128`.

### Booleans

A bool, represents either `true` or `false`. 

### Strings

Strings within the Go language are what we would call character slices.


### Constants

Constants are our final basic data type within the Go language. They allow us to specify immutable values that will not change throughout the course of our programs execution.