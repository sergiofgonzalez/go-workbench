# Mini-project: Hello, reading files!

The goal of this mini-project is to illustrate how to read a file.

The `os` package provides a `ReadFile` method defined as:

```go
func ReadFile(name string) ([]byte, error)
```

This function is quite low-level and return the file contents as a slice of bytes.

In many scenarios, the result of `ReadFile` should be transformed into a a slice of strings for managing its contents. The package `strings` contains methods such as `Split`, `SplitAfterN`, `SplitN`, `Cut`, and `Fields` that help you do that.

In particular, in this example we read a text file with lines of text and use.

Alternatively, you can open the file using `os.Open()` and then use the file pointer returned by that function to invoke `Read` in chunks.

Also, you can use `ReadLine` which is also pretty low-level.

Finally, the use of `bufio.NewScanner`, `Split`, and `Scan` is demonstrated.