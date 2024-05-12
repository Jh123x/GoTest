## Array Asserts

Array asserts are used to check if an array contains a value, does not contain a value, contains all the values, and if two arrays are equal.
To find out more about how the asserts can be used, you can take a look at the unit tests [here](https://github.com/Jh123x/GoTest/blob/main/assert/arrays_test.go "Array Test").

### Contains

Contains checks if the array contains the given value.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestContains(t *testing.T) {
    arr := []int{1, 2, 3, 4, 5}
    assert.Contains(t, arr, 3, "3 should be in the array") // Passes as 3 is in the array
    assert.Contains(t, arr, 6, "6 should be in the array") // Fails and errors as 6 is not in the array
}
```

### NotContains

NotContains checks if the array does not contain the given value.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestNotContains(t *testing.T) {
    arr := []int{1, 2, 3, 4, 5}
    assert.NotContains(t, arr, 6, "6 should be in the array") // Passes as 6 is not in the array
    assert.NotContains(t, arr, 3, "3 should be in the array") // Fails and errors as 3 is in the array
}
```

### ContainsAll

ContainsAll checks if the array contains all the given values.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestContainsAll(t *testing.T) {
    arr := []int{1, 2, 3, 4, 5}
    assert.ContainsAll(t, arr, []int{1, 2, 3}, "1, 2, 3 should be in the array") // Passes as 1, 2, 3 are in the array
    assert.ContainsAll(t, arr, []int{1, 2, 6}, "1, 2, 6 should be in the array") // Fails and errors as 6 is not in the array
}
```

### EqualArray

EqualArray checks if the two arrays are equal.

```go
package main

import (
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestEqualArray(t *testing.T) {
    arr1 := []int{1, 2, 3, 4, 5}
    arr2 := []int{1, 2, 3, 4, 5}
    assert.EqualArray(t, arr1, arr2, "Arrays should be equal") // Passes as the arrays are equal
    assert.EqualArray(t, arr1, []int{1, 2, 3, 4}, "Arrays should be equal") // Fails and errors as the arrays are not equal
}
```
