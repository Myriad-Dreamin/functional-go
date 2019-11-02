# functional-go
 functional-go

## Installation

```go
go get github.com/Myriad-Dreamin/functional-go
```
or just import it in your package

```go
package traits // import "github.com/Myriad-Dreamin/functional-go"
```


```go
type BaseTraits struct{ ... }
```

```go
type DecayResult struct{ ... }
```

```go
type DecayResult2 struct{ ... }
```

```go
type DecayResult3 struct{ ... }
```

```go
func NewBaseTraits(t interface{}) BaseTraits
```

```go
func Decay(d interface{}, e error) DecayResult
```

```go
func Decay2(d, d2 interface{}, e error) DecayResult2
```

```go
func Decay3(d, d2, d3 interface{}, e error) DecayResult3
```