## Pointer Asserts

### Nil

Nil checks if the pointer is nil.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest"
)

func TestIsNil(t *testing.T){
    var a *int
    gotest.IsNil(t, a, "a is not nil") // This returns true as a is nil

    var b *int
    gotest.IsNil(t, b, "b is not nil") // This returns false and errors as b is not nil
}
```

### NotNil

NotNil checks if the pointer is not nil.

```go
package main

import(
    "testing"

    "github.com/Jh123x/gotest"
)

func TestNotNil(t *testing.T){
    a := new(int)
    gotest.NotNil(t, a, "a is nil") // This returns true as a is not nil

    var b *int
    gotest.NotNil(t, b, "b is nil") // This returns false and errors as b is nil
}
```
