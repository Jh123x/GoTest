# GoTest

A faster testing library using generics

## Goals

1. Faster debugging (Different types will be flagged up during compile time)
2. Faster testing (No use of `reflect` to check types)

## Usage

```go
package main

import (
    "fmt"
    "testing"

    "github.com/Jh123x/gotest"
)

func TestAdd(t *testing.T) {
    gotest.Test(t, func(t *gotest.T) {
        assert.Equal(Add(1, 2), 3)
        assert.Equal(Add(1, 2), 4)
    })
}
```

## Installation

```bash
go get github.com/Jh123x/gotest
```
