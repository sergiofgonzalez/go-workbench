# append
> grokking the append function by building one

## v0: `appendInt(x[]int, y) []int`

In this first version, we create a function that appends a single element to a given slice of ints.

## v1: `appendInt(x[]int, y ...int) []int`

In this second version, we allow passing one, several, or even a whole slice to the `appendInt` function.