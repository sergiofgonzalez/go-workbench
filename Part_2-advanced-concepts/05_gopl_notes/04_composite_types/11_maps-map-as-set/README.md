# dedup
> illustrate how to use a map to implement a set

Write a program that reads a sequence of lines from the standard input and prints only the first occurrence of each distinct line.

## v0: initial implementation

We know from a previous exercise that the `bufio` package features a `Scanner` type that reads input and break into lines or words.

```go
input := bufio.NewScanner(os.Stdin)
```

Then, each call to `input.Scan()` will read the next line and remove the newline character from the end. It will also return `true` if there's a line or `false` otherwise, so that we can find when there's no more input.

Finally, you'll be able to access the line using `input.Text()`.

You can test it running:

```bash
go run . < dedup.go
```

You'll notice how your empty lines disappear from your program.