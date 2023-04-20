# maps using TDD
> this example illustrates some basic concepts on maps.

## v0: our first attempt
Write a simple program that features a `Search` function that takes a map and a key and returns the string value found for that key in the map.

## v1: using a custom type for the map

Using `map[string]string` in the program doesn't feel very idiomatic. Enhance the previous version defining a new type `Dictionary` which aliases `map[string]string` and implementing `Search` as a method instead.

Also, improve your tests to accommodate the situation in which you provide a word that is not found in the dictionary.

HINT: the `Search` method should be updated to return an error along with the word. That error will signal when the word is not found in the dictionary.

Improve the solution to use a sentinel error `ErrWordNotFound`, and make sure to return an empty string in that error condition.

Refactor the tests as necessary.

## v2: adding a new `Add` method

Enhance the previous version to include a new `Add` method to the `Dictionary` to be able to include new words to the dictionary.

As maps will update the corresponding value when doing `d[k]=v` for key that already exists, make sure that you take that design decision in your approach. To make it idiomatic, define a `ErrWordAlreadyExists` sentinel error.

## v3: refactoring sentinel errors

You can see in the previous approach that there's a little bit of duplication associated with the way in which we're creating sentinel errors:

```go
type wordNotFoundErr string

func (e wordNotFoundErr) Error() string {
	return string(e)
}

type wordAlreadyExistsErr string

func (e wordAlreadyExistsErr) Error() string {
	return string(e)
}

const ErrWordNotFound = wordNotFoundErr("could not find the word you were looking for")

const ErrWordAlreadyExists = wordAlreadyExistsErr("word already exists")
```

This duplication can be reduced by defining a generic `dictionaryErr` alias and using it as a constructor for our errors. As the base type for them will still be a string, the errors defined in such a way can still be used as sentinels.

## v4: adding an Update and a Delete method

In the same way as was done for `Add`, include a new `Update` and `Delete` methods that should let you update and delete existing definitions and remove existing keys. Define `ErrWordDoesNotExist` to signal the situation in which you try to update a definition for a word that is not present in the dictionary. To favor idempotency, ignore the situation in which you are removing a non-existent key from a dictionary.