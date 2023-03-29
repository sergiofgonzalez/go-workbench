# Mini-project: Wordle in Go

The goal of this mini-project is to create a varianto of Wordle in Go.

In the official game, you have to guess a word of 5 characters in 6 attempts. After each attempt, the game tells you, for every character, whether it belongs to the solution, and whether it has the correct position.

Along the way we'll learn:
+ to write a program that picks a random word in a list
+ read the guesses of the player from the standard input
+ give feedback on whether the characters are correctly placed or not
+ let the player know if they win or loses after having tried the max number of unsuccessful attempts.

## v0: setting up shop

For the very first version, we will implement a program that will have a hardcoded solution, allowing only one guess, and the program will answer ok or not ok.

1. Create a new package after the name of the game ("gordle"). Within it, we will create a directory `gordle/` that will contain the files in the gordle package. Outside you will have the `main.go` and the `go.mod` file.

2. In the `gordle/` package, create a `game.go` to hold all the information needed to play a game of gordle. In the file, create an empty structure named `Game`.

3. Create a `New()` function in `game.go` that returns a pointer to a new `Game` structure.

4. Define a method `Play()` that displays the prompt:

    ```
    "Welcome to Gordle!"
    Enter a guess:
    ```

5. In `main.go`, use `New()` and invoke the `Play()` method.

## v1: reading input from stdin

The `bufio` package has a useful method `ReadLine` on the `Reader` struct that returns a single line, not including the end-of-line.

We will use it to read the input from our user.

1. Add in our `Game` struct, a pointer to a `bufio.Reader`, and call it `reader`.

2. Update the `New()` function so that it receives an `io.Reader` argument and use it to initialize the `bufio.Reader` of the struct. Call it `playerInput`.

3. Define a constant `wordLength` and make it equal to 5. We will use it to check that the input provided by the user is of the expected length.

4. Define a method `ask()` on `Game` that will read the input from the user until a valid suggestion is made. Once made, it will be returned.

5. Implement the `ask()` method. The method should begin by showing a message to to player telling them that it expects a wordLength input. Then, you can implement a while true (for) that will keep asking the user for input until no error is returned by the `Readline()` method. If an error is found, a message should be written in stderr using `Fprintf(os.Stderr)`. Otherwise, if no error is found, you should read the input and place it on a variable `guess` as an array of Unicode points (`[]rune`). HINT: You might need to transform the `[]byte` returned by ReadLine into a string first.

6. In the same method, use `len(guess)` to check that the word typed by the user is of the expected length. Again, if it's not you should display a message of stderr and keep asking the user for input. Otherwise, if it's of the expected length, you should return it.

7. Modify the implementation of the `Play()` method, so that it invokes the `ask()` method.

7. Update `main.go` wiring the Stdin to the reader to get the input from the user and check that both the happy path and error path for ask (the loop) work. The program should gracefully terminate when you type a 5 character word, and keep asking for input when typing an invalid input. HINT: you only need to wire `os.Stdin` when constructing the the Game.

## v2: testing `ask()`

1. Create an internal test file in which we will test our `ask()` method (let's call it `game_internal_test.go`).

2. In the file, create a `TestAskHappy()` testing method. In it, provide a `strings.NewReader()` when instantiating a `Game` struct. In the asserting, check that the output of `ask()` is the expected array of runes by using the `slices.Equal()` from the slices module (HINT: you will need to do `go get slices` to add that module).

3. Once you have the first version of the `TestAskHappy()` ready, modify it to run tests including runes beyond the regular ASCII codepoints (HINT: use Spanish "cañón" and Chinese "你好世界!").

## v3: refactoring `ask()`

Let's give some structure to our `ask()` method, creating a supporting function that will check whether the user has typed a correct guess or not. For now, we will not check that the user has typed the correct word, only that the typed word is of the correct length.

1. Modify the `ask()` method so that the word length validation is taken to a a `validateGuess(guess) error` method.

2. In the `validateGuess()` method, check that the length of the guess is the same as the established `wordLength`, and return an error otherwise or `nil` when that matches. (keep reading...)

3. The `validateGuess()` function must return an `error` when the length doesn't match. In order to follow Go's conventions, we will wrap and propagate an error. In order to do so, define a variable outside of the method named `errInvalidLengthWord` of type `error` with a message informing the guess is of an incorrect length. Then, in the method, wrap the `errInvalidLengthWord` into a message with further details using `fmt.Errorf()`. As part of the returned value use `"%w"` to wrap the `errInvalidLengthWord` error.

4. Adjust the implementation of `ask()` to call and check the result of `validateGuess()`. When the validation fails, send a message to the Stderr showing the contents of the error received.


Go's take on error handling is to use functions that return an error as the last list of returned values of the function.

As a result, the first line that follows that call to such a function must be:

```go
if err != nil {
    ...
}
```

When we catch an error, the best thing we can do it handle it. If there's nothing we can do about it, it is our responsibility to propagate it to the upper layer nicely wrapped. This will provide context to that layer so that it can take the decision to handle it, or propagate it further.

The simplest way to wrap the error is to do: `fmt.Errorf("... %w ... ", err)`. The `"%w"` will wrap the error and `return fmt.Errorf()` will return the wrapped error.

5. Validate that the program works as expected using `go run .` and testing with a 5 character word (happy path) and less than/greater than 5 characters (error path).

6. In `game_internal_test.go` create a new test for `validateGuess()` function. Prepare several test cases with a word of the expected length, a word with more characters, a word with less characters, an empty word, and a nil input. As in some cases, we will be expecting `validateGuess()` to return an error, make sure that you make use of the function `errors.Is()` that lets you check if a given propagated error is of the expected type (`errInvalidWordLength` in our case).


The `errors` package provides the function `errors.Is()` which reports whether any error in the `err`'s chain matches the `target` error. This function is very handy when dealing with wrapped errors, and can simplify error reporting in our application tool, so that the message doesn't become too verbose.

7. In preparation for the subsequent step, create a new function in `game.go` named `splitToUppercaseCharacters` that takes a string and returns a slice of runes with all the characters in uppercase.

8. Create a test for the function that validates several scenarios.

9. Modify the implementation of `ask()` to make use of `splitToUppercaseCharacters()` and remove the previous `guess := rune[](string(guess))`.

## v4: check for victory

In this version, we will check if the player's input matches the solution, allowing them to try for a given number of attempts.

1. Enrich the Game structure so that it includes the `solution` as a `rune[]` and the `maxAttempts` as an int.

2. Update the `New` function, so that it receives now the solution and the maxAttempts as arguments and wire them into the `Game` structure. HINT: In order to simplify the management, define the solution as a `string` and apply `splitToUppercaseCharacters` to the argument received.

3. Modify the `play` method so that after showing the initial message "Welcome To Gordle!" then enters a loop for `maxAttempts`. If the `guess` taken from the user is equal to the solution, a message informing the use that they have guessed the word in x number of guesses and that the word was the solution should be printed. If after `maxAttempts` the user has not guessed the solution, show a message informing the user they've lost the game and what the solution was. HINT: make the messages attractive using emojis (Win-. in Windows).

4. Remove the `wordLength` constant and refactor the code accordingly.

5. Update the tests to validate they're still working.

6. Update the `main.go` file so that you can start playing. For now, set the word as a hardcoded value.

## ToDo
