# Mini-project: Sequence type implementing multiple interfaces
> an exploration of types and interface implementation

## Interfaces

A type can implement multiple interfaces. As an illustration of this concept, define a type `Sequence` as an ordered collection of ints (`[]int`).

Then implement the `sort.Interface` which requires defining:
+ `Len`
+ `Less`
+ `Swap`

Then, define also the `Copy` method that returns a copy of the sequence.

Finally, make `Sequence` implement also the *Stringer* interface, and code the implementation in a way that shows the results sorted in ascending order:

```go
s := Sequence{4, 1, 2, 5, 3}
fmt.Println(s)  // [1 2 3 4 5]
```