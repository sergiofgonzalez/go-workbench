# Mini-project: A polyglot hello, world!

## v0

Write a variation of the "Hello, world" program that is capable of handling greeting you in English, Spanish, French and German.

The program must do so by delegating the reception of the greeting in a function `greet(l language)`. Where language should be a type alias of string.

When the greet function receivs a language that it doesn't support it should return an empty string. Take a simple approach using a `switch` that checks the received language.

Implement the necessary tests in your program.

## v1

Enhanced the prior version by defining a map in the program so that instead of using we access the map and returning the corresponding greeting message if found, or the string `unsupported language: "<lang-code>"`.

Fix your tests to take into account the changes.

## v2

Improve the implementation of the tests by implementing a slight variation:

1. Declare a `TestGreet(t *testing.T)` function that will take care of both the happy and error path for the testing of `greet()`.

2. Define a new type `TestCase` within the function. This struct will represent the individual test cases the function will go through. As a result, it will need two fields, the language and the expectedGreeting.

3. Define a map named `tests` that identify the different tests to run identified by a string title for easier maintenance and report (i.e., an array could have been used here, but a map will help us with a meaningful labelling of tests).

4. Run the tests by going over the scenarios described by the `tests` map. Explore the `t.Run()` function that lets you run a test in a *more programmatic* fashion by passing an anonymous function to be run.


## v3 use the `flag` package to read from the terminal

Go provides support for parsing the command line arguments in both the `os` and `flag` packages. The former is quite low-level so in the majority of the cases, you will prefer the ussage of `flag` which will give you support for:
+ Parsing command-line arguments
+ Convert them to the right type

Improve the implementation of the program by exposing a parameter on our command line executable so that the end-user can choose the language.

1. Expose the parameter "lang" using the `flag.StringVar()`  function. Use "en" as the default language code.

2. Parse the incoming command-line arguments using `flag.Parse()`.

3. Confirm that you can start passing arguments from the CLI, and that when you don't you get the message in the default language.

| HINT: |
| :---- |
| You can use `go run .` for the default, and `go run . --lang=es` for Spanish. |

4. Remember to fix your test.