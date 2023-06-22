# comma
> illustrating strings and bytes

Create a program that takes a string representation of an integer and introduces a `','` separator every three places:

+ `comma("12345")`   -> `12,345`
+ `comma("1234567")` -> `1,234,567`

## v0: initial implementation

The solution is trickier than it seems, and a recursive function works best. You should:
+ return the received string as-is, if its length is less than or equal to 3.
+ call recursively with a substring consisting of all but the last three characters, and append a comma, and the last three characters of the result.

## v1: using string builder and an iterative approach

An iterative take, which might be flaky.