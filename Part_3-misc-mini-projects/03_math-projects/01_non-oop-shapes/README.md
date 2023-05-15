# Playing with shapes without OOP in Go

## v0: Creating a rectangle struct with some methods

Start by creating a data structure defining a rectangle, which will be identified by its width and height.

Define the following methods:
+ a constructor that would let you create a new `Rectangle`
+ an `area` method
+ a `scale` method that  returns a new rectangle with its width and height scaled by the given number
+ an `equal` method that returns true if two rectangles aer equal.
+ a `String` method implementing the Stringer interface that returns the String "Rectangle {width} by {height}".
+ a `Square` function that returns a rectangle whose both sides are of the same size.


## v1: Creating a square struct

Now we need to create a square struct with the same methods as the rectangle, but we want it to have its own type.

As there is no inheritance concept in Go, we'll need to find another solution to avoid a massive code duplication. This can be done using composition, as can be seen in this version.

## v2: Abstracting with a Shape interface

In this version we try to create something similar to an abstract class, so that we can define functions that work on an abstraction of a Shape (whether it is Rectangle or a Square).

To do so, we create a `Shape` interface and put the common methods there:

```go
type Shape interface {
	Area() float64
	Scale(factor float64) Shape
	Equal(other Shape) bool
}
```

Note that this will trigger some changes in the `rectangle` and `square` packages to accommodate some of the methods to the interface (they will need to be aware of this new `Shape` interface).

In particular:
+ `Scale` signature must be changed to return a `Shape`. Note that the implementation won't have to be changed though.

    ```go
    func (r *Rectangle) Scale(factor float64) shapes.Shape {
	    return New(r.Width * factor, r.Height * factor)
    }
    ```

    Note how `shapes.Shape` and the `*Rectangle` returned by `New` are compatible. One of the reasons you don't need to have pointers to interfaces in Go.

+ `Equal` signature and implementation will have to be changed. In this case, doesn't have to do with the return type, but rather, within the implementation, we need to make sure that when comparing shapes of a different type we get false.

    In theory this can be done with a simple type assertion, but there's a catch &mdash; as all the methods are defined on pointer receivers, the type assertion will look like:

    ```go
    rOther, ok := other.(*Rectangle)
    ```

    So that the implementation is:
    ```go
    func (r *Rectangle) Equal(other shapes.Shape) bool {
	    rOther, ok := other.(*Rectangle)
	    if !ok {
		    return false
	    }

	    return r.Width == rOther.Width && r.Height == rOther.Height
    }
    ```
