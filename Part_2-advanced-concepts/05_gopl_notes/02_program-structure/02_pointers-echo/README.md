# echo
> playing with pointers and the `flag` package

Create an implementation of the `echo` command that outputs in the standard output the arguments that are passed to it.

It should allow you to pass to optional flags that tailor certain behavior of the tool:
+ `-n` : causes echo to omit the trailing newline that is added by default
+ `-s sep` : causes echo to use a custom string to separate the different arguments it receives.

HINT: rely on the `flag` package for this functionality instead of accessing `os.Args` directly. You might need to call `flag.Parse` and `flag.Args` to parse and get the list of arguments received.

To test it type:

```bash
$ go run . hello world again
hello world again

$ go run . -n hello world again
hello world again $

$ go run . -sep "#" hello world again
hello#world#again

$ go run . -sep "#" -n hello world again
hello#world#again $
```