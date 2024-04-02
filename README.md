# GoTest

A faster testing library using generics

## Installation

```bash
go get github.com/Jh123x/gotest
```

## Usage

```go
package main

import (
    "fmt"
    "testing"

    "github.com/Jh123x/gotest/assert"
)

func TestAdd(t *testing.T) {
    assert.Equal(1+ 2, 3, "1 + 2 should be 3")
    assert.Equal(1+ 2, 4, "1 + 2 should be 4")
}
```

## Documentation

To view the full documentation, visit the docs [here](./docs/index.md "Go Test Documentation")

## Goals

1. Faster debugging (Different types will be flagged up during compile time)
2. Faster testing (No use of `reflect` to check types)
