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

## structs
like an object in js
```go
type myStruct struct {
    chungus string
    int
}

// initialize structs like anythihng else:
var chungus myStruct{chungus:"bungus", 4}
var chungus myStruct{"bungus", 4}

// you can also do
var myStruct = struct{
    chungus string
}{"bungus"}
```

### methods
structs can  hold functions like class methods
```go
func (e myStruct) someFunc() { }
```

### interface
```go
type myInterface interface {
    myFunct()
}
```

## pointers
```go
var myPtr *int // is nil, doesnt allocate
var myPtr *int = new(int) // allocates memory for an int
```

get value of pointer (dereferencing)
```go
*myPtr
```
get addy (referencing)
```go
   myVar := 42
   ptr := &myVar
```

```go
*dereference
&reference
```

functions always copy params in memory
```go
param := 1
fmt.Println(param) // some memory addy

func chungus(p int) {
    fmt.Println(p) // some other memory addy
}
```
circumvent with references
```go
param := 1
fmt.Println(param) // some memory addy

func chungus(*p int) {
    fmt.Println(p) // same memory addy
}
```

## Go multi-core concurrency

```go
go someFunc() // add "go" before a function call

```

wait for concurrency to finish
```go
import "sync"
wg := sync.WaitGroup{}

wg.Add(1) // add whenever we spawn a go routine
go someFunc() 

wg.Wait() // block code execution until finish

...
wg.Done()

```

handle transactions
```go
import "sync"
m := sync.Mutex{} // (mutual exclusion)
m.Lock() // go routines will reach this and check to see if this section has been locked by another go routine and stop/continue accordingly
// write data
m.Unlock()
```

full example
```go
import "sync"

func main() {
    m := sync.Mutex{}       
    wg := sync.WaitGroup{} 

    data := map[string]string{"chungus": "bungus"}

    for k, v := range data {
        wg.Add(1)             // Add 1 goroutine to WaitGroup counter
        go concatData(k, &data, &m, &wg)
    }

    wg.Wait()  // Wait for all goroutines
}

func concatData(key string, data *map[string]string, m *sync.Mutex, wg *sync.WaitGroup) { // referencing because otherwise each parameter would become a copy and locking/unlocking would happen to copies
    defer wg.Done()     // Ensure WaitGroup counter is decremented
    defer m.Unlock()    // Ensure mutex is unlocked
    
    m.Lock()            
    (*data)[key] = (*data)[key] = + " example" // awkward accessor bc of operator precedence ([] comes before * so data is accessed before specifying dereference)
}
```

theres also `sync.RWMutex{}` which additionally provide `RLock()` & `RUnlock()`

## Channels
- hold data
- thread safety (read and write racing)
- can be listened to

```go
c := make(chan int)
c <- 42 // assign or retrieve value


```

THERES MORE


## Generics

```go
func genericFunc[T int | string](someVar T) T{
    var idk T
    idk = someVar

    return idk
}

genericFunc[string]("chungus")

```
```go
type myStruct [T int | string]struct {
    myAttr string
    myOtherAttr T
}

```

theres also any, but for any to work whatever operation you do on it must be compatible with any type


