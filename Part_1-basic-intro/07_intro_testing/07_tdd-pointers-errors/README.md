# pointers and errors using TDD
> this example illustrates some basic concepts on pointers and errors.

## v0: using values as receivers

Write a simple program that features a `Wallet` structure that includes a `balance` field of type `int`. Then, define a couple of methods `Deposit` that increments the balance, and `Balance` which returns the actual balance.

Do not use pointers as receivers.

Confirm that the `Deposit` method does not work as expected because the method receives a copy of the original variable. Use extra logging and the syntax `&wallet.balance` to print out the address of the balance field in the test method and the `Deposit` method and validate that those are different.

## v1: using pointers as receivers

Refactor the v0 to use pointers as receivers and confirm that the test passes. Make sure to update both methods for consistency.

## v2: using type aliases

In v1, the `Wallet` struct define the balance as a plain int. It is a good practice to create a type alias to give a better indication of what the struct field represents.

Refactor v1 to create a type alias for the type of balance field, and call it `Bitcoin`. Also, update the methods so that they receive and return `Bitcoin`.

Update the tests accordingly.

## v3: implementing interfaces on types

Update the project to make `Bitcoin` implement the `Stringer` interface.

Confirm that it works as expected by breaking the test and seeing the implementation of the `String` method is invoked.

HINT: You may need to change the format specifier for your error messages from "%d" to "%s".

## v4: adding a Withdraw method

Enhance the previous approach by defining a new `Withdraw` method.

Refactor the code and tests as required to minimize duplication.

## v5: controlling overdraft

The previous version allows you to withdraw more money than you actually have in your account.

It is idiomatic in Go to signal that kind of situations by way of returning an error in the method or function.

Introduce this new functionality, so that the `Withdraw` function signals the overdraft, while leaving the state of your wallet as it was before the attempt to withdraw more funds than the available balance.

Signal the error situation with an error string informing that there are no sufficient funds to complete the withdraw. In the tests, check that you get an error, and that the error message matches your expectations.

HINT: you might want to use `err.Error()` to get a string representation of the error in an overdraft situation.

## v6: using sentinel errors

While perfectly functional, it is considered idiomatic in Go to define *sentinel errors* when you want to allow the code using a library to check for specific errors using constructs such as:

```go
err := wallet.Withdraw(10)
if err == errInsufficientFunds {
  ...
}
```

Enhance the previous version by adding a sentinel error for the insufficient funds situation.

Refactor the tests and code as necessary.