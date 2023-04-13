# Mini-project: ByteSlice
> a custom implementation of a slice of bytes implementing the `io.Writer` interface

The goal of this mini-project is to exercise a few basic Go techniques by way of implementing a custom type `ByteSlice` (namely an alias to `[]byte`), and then improve it in several steps so that it ends up implementing the `io.Writer` interface, and therefore can be used where a native `io.Writer` would be used such as in `fmt.Fprint`.

## v0: Implementing an `Append` method

1. Define a type `ByteSlice` as an alias to `[]byte`.

2. Define a method `Append(data []byte) []byte` using a ByteSlice value as a receiver.

3. Implement the function so that it appends the elements of data to the ByteSlice value into a new ByteSlice that is then returned.

4. Implement a set of tests for the `Append` method.

## v1: Improving the `Append` method using pointers

The previous version of the function is a little clumsy, as it requires a ByteSlice value but returns a new []byte. In this version, you have to improve the implementation so that it updates the method receiver.

1. Change the method signature, so that it receives pointer to a ByteSlice and not a value. Also, the function does not return anything now.

2. Change the implementation of the function so that:
  + at the beginning of the function, you dereference the pointer receiver to get direct access to the ByteSlice contents.
  + at the end of the function, you copy the resulting of the appended ByteSlice (the one you returned in v0), into the contents of the pointer receiver.

3. Update the tests, so that they work with the new method signature.

## v2: Implementing the `io.Writer` interface

The `io.Writer` interface is defined as follows:

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

In this version, you have to change the `Append` function once more so that it satisfies the `io.Writer` interface.

In order to confirm it works where any `io.Writer` would work, use a ByteSlice within an `Fprintf` statement.