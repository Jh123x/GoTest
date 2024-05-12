## Pointer Asserts

Pointer asserts are used to check if a pointer is `nil` or not.
To find out more about how the asserts can be used, you can take a look at the unit tests [here](https://github.com/Jh123x/GoTest/blob/main/assert/pointers_test.go "Pointer Unit Test").

### Nil

Nil checks if the pointer is nil.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestIsNil(t *testing.T){
    var a *int
    assert.IsNil(t, a, "a is not nil") // This returns true as a is nil

    var b *int
    assert.IsNil(t, b, "b is not nil") // This returns false and errors as b is not nil
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
    assert.NotNil(t, a, "a is nil") // This returns true as a is not nil

    var b *int
    assert.NotNil(t, b, "b is nil") // This returns false and errors as b is nil
}
```
