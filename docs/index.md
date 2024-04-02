## Go Test Documentation

A faster testing library using generics

### Goals

1. Faster debugging (Different types will be flagged up during compile time)
2. Faster testing (No use of `reflect` to check types)

## Functions

### Compare Asserts

- [Equal](./compare.md#Equal)
- [NotEqual](./compare.md#NotEqual)
- [EqualMap](./compare.md#EqualMap)
- [Greater](./compare.md#Greater)
- [Less](./compare.md#Less)
- [GreaterOrEqual](./compare.md#GreaterOrEqual)
- [LessOrEqual](./compare.md#LessOrEqual)

### Pointer Asserts

- [Nil](./pointer.md#Nil)
- [NotNil](./pointer.md#NotNil)

### Bool Asserts

- [True](./bool.md#True)
- [False](./bool.md#False)

### Array Asserts

- [Contains](./arrays.md#Contains)
- [NotContains](./arrays.md#NotContains)
- [ContainsAll](./arrays.md#ContainsAll)
- [EqualArray](./arrays.md#EqualArray)
