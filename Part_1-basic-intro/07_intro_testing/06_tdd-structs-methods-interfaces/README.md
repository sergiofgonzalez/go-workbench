# structs, methods, and interfaces using TDD
> this example illustrates some basic concepts on structs, methods, and itnerfaces using TDD.

## v0: writing functions

Using TDD, write a couple of functions `Perimeter` and `Area` that calculate the perimeter and area of a rectangle. The functions should receive the width and height of the rectangle as floating-point values.

## v1: using structs

While v0 is perfectly functional, the current implementation does not provide any information related to the specific shape we're dealing with (rectangles).

Refactor the solution so that everything is based on a Rectangle shape defined as a struct.

## v2: using methods

Refactor v1 solution so that it can be applied to different shapes, specifically rectangles and circles.

As Go does not support function overloading, define `Area` and `Perimeter` functions as methods rather than standalone functions.

## v3: using interfaces

While v2 version is fully functional, we miss the possibility from OOP of being able to create a super class and then a couple of subclasses for Rectangles and Circles.

The way of getting similar capabilities in Go is by way of using interfaces.

Begin by creating an interface `Shape` with a couple of functions `Perimeter` and `Area`.

Then, start by refactoring your tests so that you reduce the duplication by defining helper functions `checkArea` and `checkPerimeter` that invoke the corresponding method on the shape, and return whether the expected result matches what the method returns.

## v4: further refactoring on tests

Perform further refactoring on the tests to make use of *"table-driven tests"*, that is, a map of tests that describes all the scenarios to tests.

Make use of the recently defined `Shape`.

## v5: introducing new shapes

Enhance the previous version to include a new shape: `Triangle`.


| NOTE: |
| :---- |
| You might be tempted to go one step further and define methods for `Shape`, but Go does not allow that. Any named type can be used as a receiver when writing a method, except for a pointer or an interface. |