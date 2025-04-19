använde AI för
- connection string
- dockerfile run multiline string syntax
- json marshalling

# Learn go

## typing

data type size will default to system architecture so 32 bit systems will translate int to int32 for example

uints are positive

```go
// strings are immutable
chungus := ""
chungus = "bungus" // creates a new string
chungus += "bobungus" // creates a new string

// instead, do this:
import "strings"
var strBuilder strings.Builder // creates a slice in the background
strBuilder.WriteString("bungus") // append()
strBuilder.WriteString("bobungus")
chungus := strBuilder.String()

```

import "unicode/utf8"
utf8.RuneCountInString() to get len of string

go types are 0 or false by default

```go
func someFunc() (multiple, return, vars) { }

var1, var2, var3 = someFunc()
```

```go
import "errors"
err = errors.New()

err.Error() // error string
```

# collections

#### ordered collections
- array
- slice
- string

#### unordered collections
- map
- set

## arrays
arrays are fixed length
```go
var intarray [5]int
intarray := [5]int{ 1, 2, 3 }
intarray := [...]int{ 1, 2, 3 } // inferred length
```

## slices
slices wrap arrays with more functions
```go
intSlice := []int{ 1, 2, 3 } // omit length and u get a slice
intSlice := make(int, 3, 6) // another way to initialize a slice

intSlice = append(intSlice, 4)
intSlice = append(intSlice, someArray...) // spread

cap(intSlice) // capacity of slice
```
you can be performant by minimizing capacity based on needs, 
because you minimize how much you have to reallocate new contiguous arrays

## maps
```go
    stringIntMap := map[string]int{ "chungus":1, "bungus":2 }
    make(map[typeOfKey]typeOfVal)

    stringIntMap["chungus"] -> 1
```

accessing map values returns two values
```go
    stringIntMap["bobungus"] -> 0 (default value of type)
    bobungus, exists := stringIntMap["bobungus"] // exists is true/false depending on if the value exists in the map
```

`delete(map, key)` delete by reference

`for k, v := range array|slice|map {}`

