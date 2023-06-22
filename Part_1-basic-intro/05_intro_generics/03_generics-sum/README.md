# sum
> learning generics through an example

## v0: creating different functions

In this version we create three variants of the same function that sums the numbers received.

## v1: using generics

In this version, we just create a single function using generics, and the types are explicitly listed in the function definition:

```go
func Sum[T int | int64 | float64](nums []T) T {
  ...
}
```

## v2: using generics and a type

In this version, we create a type definition that includes the diff

```go
func Sum[T int | int64 | float64](nums []T) T {
  ...
}
```
