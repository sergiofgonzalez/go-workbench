# Reflection in Go using TDD
> this example illustrates how to use reflection in Go

| NOTE: |
| :---- |
| While I tried to include as much comments as possible, I got to a point where it became really complicated to follow, and decided to just copy the code understanding it, rather than describing what shoould be done and try to do it myself.<br>Feel free to re-read the [reflection](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reflection) chapter again if you get lost. |


In this example you will implement a solution for the following challenge:

> write a function `walk(x interface{}, fn func(string))` which takes a struct `x` and calls `fn` for all string fields found inside.


## v0: first implementation

Start by creating a test for the `walk` function. As part of the assertion, you can check that the fields traversed by `walk` are the ones you expect.

That is, for the test you can do something like:
1. Define the struct `x` as simple struct with a single string field `Name` initialized with a certain string value (e.g., "Jason Isaacs").

2. Declare a slice `got` that will collect the string fields identified by the `walk` function.

3. Invoke `walk` passing as a function one that simple appends to `got` the identified fields.

4. Assert that `got` has a single element, whose content is `"Jason Isaacs"`.

You can then provide a dummy implementation of `walk` that simply calls the function with a hardcoded value.


## v1: implementing `walk` using reflection


For the implementation of the `walk` function, you will need to use reflection. For the first implementation you can do something like:

```go
val := reflect.ValueOf(x)		// get x as a Value using reflection
field := val.Field(0)				// naively access the first Field in the Value
fn(field.String())					// use `String` to get the underlying value of the Field
```

Confirm that now the tests run correctly.

## v2: Detecting the number of fields

We'll need to test several scenarios, and therefore, it makes sense to start preparing `TestWalk` to use table-based tests.

Once you have accommodated the table-based structure, you can include another test with a struct with two string fields.

The previous version assumed that everything tht is passed to walk is a struct with a single string field.

You will have to improve the implementation of `walk` so that it is able to handle multiple fields.

HINT: you will have to implement a loop using `Value.NumField()`.

## v3: Detecting different types of fields

The previous version is still very buggy, as it assumes that whatever you're passing to `walk` is a struct with zero or more string fields.

Let's start making it more robust by detecting the different types you might find in the struct.

Write a test that checks that you can pass a struct with a string and an int, and then implement a `walk` function that makes the test go green.

## v4: Handling recursion

The previous version is much more robust as it handles different types of fields and can handle more than one field.

We can start by writing the test first, by declaring a struct with nested fields such as the following:

```go
struct {
	Name string
	Profile struct {
		Age  int
		City string
	}
}
```

This will require a more contrived approach in the implementation of `walk` requiring recursion. Also, you will need to use `Field.Interface()` to get the value of the field as an `interface{}`.

## v5: refactoring with a switch and handling pointers

The first thing you can do is do a bit of refactoring of the `walk` function to use a `switch` if you haven't used it already.

Then, you can add another test case that includes a pointer to a struct, rather than a struct. Initially, it should fail with a panic because you're trying to use `NumField` on a pointer.

HINT: check if the received element is a pointer by comparing its kind to `reflect.Ptr`, and if so, use `Value.Elem` to dereference it.

Finally, refactor the `walk` implementation a little bit, creating a function `getValue` which will also be helpful afterwards.

## v6: handling slices

Write now we're only handling structs, but we should enable `walk` to handle slices too!

## v7: refactor

While the code of the walk function works, it's quite complicated to understand. It will be much better to refactor using `switch`, checking what's the type of value received, and acting accordingly.

## v8: extreme refactoring

There's some additional refactoring that can be done to make the implementation of the `walk` function a bit more succinct and reduce duplication.

## v9: handling arrays and maps

In this version we will support arrays, which should be trivial now that we have the support of slices in place.

Note that you can support multiple conditions in switch with the syntax:

```go
switch var {
case cond1, cond2:
	// .. do something when cond1 or cond2 is true ...
}
```

Confirm that it now works for arrays and maps.

## v10: rolling back one of our refactoring activities

Adding support for maps will make one of our recent refactoring less appealing, because we have to iterate through the map keys.

As a result, it is much better to bring back the iteration back to make it more consistent in most of the `val.Kind()` cases.

Additionally, maps in Go don't guarantee that the order is preserved, so the `slice.Equals` we're using is bound to fail.

One way to fix it is to use `slice.Contains` instead of `slice.Equals`, but this will require getting the test for maps outside of the existing table.

## v11: adding support for channels and functions

The last types we need to support are channels and functions (those are first-class citizens in Go).

You can implement it for the values that are waiting in the channel, and for the values returned by a function.

