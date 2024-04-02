## Comparison Asserts

### Equal

Equal checks if the two values are equal.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestEqual(t *testing.T) {
    assert.Equal(t, 1, 1, "1 should be equal to 1") // Passes as 1 is equal to 1
    assert.Equal(t, 1, 2, "1 should be equal to 2") // Fails and errors as 1 is not equal to 2
}
```

### NotEqual

NotEqual checks if the two values are not equal.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestNotEqual(t *testing.T) {
    assert.NotEqual(t, 1, 2, "1 should not be equal to 2") // Passes as 1 is not equal to 2
    assert.NotEqual(t, 1, 1, "1 should not be equal to 1") // Fails and errors as 1 is equal to 1
}
```

### EqualMap

EqualMap checks if the two maps are equal.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestEqualMap(t *testing.T) {
    a := map[string]int{"a": 1, "b": 2}
    b := map[string]int{"a": 1, "b": 2}
    assert.EqualMap(t, a, b, "a should be equal to b") // Passes as a is equal to b
    c := map[string]int{"a": 1, "b": 2}
    d := map[string]int{"a": 1, "b": 3}
    assert.EqualMap(t, c, d, "a should be equal to b") // Fails and errors as c is not equal to d
}
```

### Greater

Greater checks if the first value is greater than the second value.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestGreater(t *testing.T) {
    assert.Greater(t, 2, 1, "2 should be greater than 1") // Passes as 2 is greater than 1
    assert.Greater(t, 1, 2, "1 should be greater than 2") // Fails and errors as 1 is not greater than 2
}
```

### Less

Less checks if the first value is less than the second value.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestLess(t *testing.T) {
    assert.Less(t, 1, 2, "1 should be less than 2") // Passes as 1 is less than 2
    assert.Less(t, 2, 1, "2 should be less than 1") // Fails and errors as 2 is not less than 1
}
```

### GreaterOrEqual

GreaterOrEqual checks if the first value is greater than or equal to the second value.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestGreaterOrEqual(t *testing.T) {
    assert.GreaterOrEqual(t, 2, 1, "2 should be greater than or equal to 1") // Passes as 2 is greater than 1
    assert.GreaterOrEqual(t, 1, 1, "1 should be greater than or equal to 1") // Passes as 1 is equal to 1
    assert.GreaterOrEqual(t, 1, 2, "1 should be greater than or equal to 2") // Fails and errors as 1 is not greater than 2
}
```

### LessOrEqual

LessOrEqual checks if the first value is less than or equal to the second value.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestLessOrEqual(t *testing.T) {
    assert.LessOrEqual(t, 1, 2, "1 should be less than or equal to 2") // Passes as 1 is less than 2
    assert.LessOrEqual(t, 2, 2, "2 should be less than or equal to 2") // Passes as 2 is equal to 2
    assert.LessOrEqual(t, 2, 1, "2 should be less than or equal to 1") // Fails and errors as 2 is not less than 1
}
```
