## Bool Asserts

Bool asserts are used to check if a value is `True` or `False`.
To find out more about how the asserts can be used, you can take a look at the unit tests [here](https://github.com/Jh123x/GoTest/blob/main/assert/bool_test.go "Bool Test").

### True

True asserts that the value is `True`.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestTrue(t *testing.T) {
    assert.True(true, "should be true") // Passes as the value is true
    assert.True(false, "should be true") // Fails and errors as the value is false
}
```

### False

False asserts that the value is `False`.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestFalse(t *testing.T) {
    assert.False(false, "should be false") // Passes as the value is false
    assert.False(true, "should be false") // Fails and errors as the value is true
}
```
