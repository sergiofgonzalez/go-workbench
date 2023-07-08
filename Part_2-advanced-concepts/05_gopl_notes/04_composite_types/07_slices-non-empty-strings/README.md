# nonempty
> practicing in-place slice techniques

Write a function `nonempty(strings []string) []string` that given a list of strings, rturns the non-empty ones.


## v0: initial implementation

We use an algorithm that performs an in-place technique, reusing the slice we received as an argument.

## v1: using append

In this implementation we use append.