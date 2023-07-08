# charcount
> grokking maps

Write a program that counts the occurrences of each distinct Unicode code point in its input.

## v0: initial naive implementation

In this first version, we approach the problem without using any of the unicode libraries.

## v1: a more robust implementation

In this implementation we use `bufio.NewReader`.