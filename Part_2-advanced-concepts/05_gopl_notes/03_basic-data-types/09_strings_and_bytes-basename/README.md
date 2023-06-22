# basename
> illustrating the use of strings and bytes

Inspired by the `basename` Unix shell utility, create a function `basename(s string)` that removes any prefix of `s` that looks like a file system path with components separated by slashes, and remove any suffix that looks like a file type so that:

+ `basename("a/b/c.go")` -> `"c"`
+ `basename("c.d.go")` -> `"c.d"`
+ `basename("abc")` -> `"abc"`

## v0: using no libraries

Implement `basename` without the use of any help libraries such as `bytes` or `strings`.

## v1: reimplement using help libraries

Reimplement the previous function using Go libraries.