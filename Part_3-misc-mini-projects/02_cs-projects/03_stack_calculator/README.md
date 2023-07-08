# Stack Calculator
> Postfix evaluation of integer expressions

## Description

This exercise illustrates how to evaluate a postfix expression involving integer values such as:

```
3 6 2 - +
```

using the stack developed in [Stack ArrayStack](../02_stack-arraystack/README.md)

Additionally, in order not to copy the code, the dependency to the stack is added to the project.

### Working with the Stack from the Stack Calculator

The Stack Calculator declared the module `example.com/arraystack`.

In order to work with that module without having to copy the code we have to:

1. Find the path to the arraystack module:

    ```bash
    $ cat ../../02_stack-arraystack/arraystack-v0/go.mod
    module example.com/arraystack
    ```

2. Edit our current module to add the dependency to this module:

    ```bash
    go mod edit \
     --replace \
     example.com/arraystack=../../02_stack-arraystack/arraystack-v0/
    ```

3. Now we need to `get` the module:

    ```bash
    go get example.com/arraystack
    ```

4. And we're ready to start coding:

    ```go
    package main

    import (
      "fmt"

      "example.com/arraystack/stack"
    )

    func main() {
      s := stack.New()
      fmt.Println(*s)
    }
    ```

### StackCalc v0

As the current implementation of the Stack is not very flexible (it can only hold 10 ints), it doesn't make sense to spend a lot of time adding fancy things, so we'll stick to the basics:

+ Read the arguments from the command-line
+ Scan them to check if those are ints or operations
    + if are ints, push them to the stack
    + if are operation, pop one or two elems (depending on whether it's unary or binary op), and push the result to the stack
    + if the arg is not an int or an op, panic and exit
    + when the args are exhausted, print the result.


### StackCalc v1

In this version we switch from the unflexible array implementation, to the generic one.

Note to test it, run:

```bash
go run . 1 2 +
```