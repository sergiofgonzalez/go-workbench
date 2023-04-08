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

## v5: providing feedback to the user

The game, in its current installment, is very difficult to win. Our task is to make user's life easier by letting them know which characters of that word are in the correct position, which are in the wrong position, and which simply don't appear in the solution.

The feedback will be a list of indications that can have three values: correct position, misplaced but found in the word, and absent.

1. Create a file `hint.go` in the `gordle` package in which you define a type `hint` as a byte.

2. In the same file, define the enumeration `absentCharacter`, `wrongPosition`, and `correctPosition`.

3. Implement the `Stringer` interface for the `hint` type. Within the method, you will return ⬜️ if the character is absent from the word, 🟡 if a character is found in the word, but in the wrong position, 💚 is the character is found in the correct position, 💔 if none of the previous one is matched (which shouldn't happen).

4. Define a new type `feedback` as a slice of the `hint` type recently defined.

5. Define a method `String` on the feedback type that takes nothing and return a string in which each character is the string representation of the individual hint. Use the ``strings.Builder` type to do so to minimize the number of memory allocations.

6. Write a test for the `feedback.String()` method.

## v6: Checking the guess algorithm

1. Define a function `computeFeedback` that takes a guess and a solution and returns a `feedback` in the `game.go` file. Make guess and solution a `[]rune`.


2. Study the following algorithm which will be implemented next (this one in pseudo-code):

        ```
        for all characters of the guess {
            mark character as absent
        }

        for each character of the attempt {
            if the character in solution and guess are the same
                mark character as seen in the solution
                mark character with correct position
        }

        for each character of the guess {
            if current character already has a hint {
                skip to the next character
            }

            if character is in the solution and not yet seen {
                mark character as seen in the solution
                mark character with correct position status
            }

        }

        return the hints
        ```

3. Implement the function by using a slice of hints to mark characters of the guess with their appropriate hint, and a slice of booleans to mark the characters of the solution as either seen or not yet seen.

4. Implement a test for the `computeFeedback` function. Hint: Define an `Equal` method in `hint.go` that determines that two feedbacks are equal. While you can use `slices.Equal` implement it yourself there.

5. Integrate the `computeFeedback` invocation in the `Play` function and display the feedback right after that.

## v7: adding random words

1. Create a directory `corpus/` with a file names `english.txt`. This file will contain a list of uppercase English words, one per line.

2. Create a new file `corpus.go` where all the method releated to the solution words will be added.

3. Define an error `ErrCorpusIsEmpty` to signal when the file is empty. To do it correctly, start by defining a type `corpusError` that is an alias to string. Then, define the `Error()` method on that type. The implementation of the method can simply return the string of the error value.

4. Define a function `ReadCorpus`. The function will take the path to the file, and will return a slice of strings and an error.

5. Implement the function as below:
    + Read the file contents using `os.ReadFile`.
    + if an error is found, return the appropriate error (HINT: use `fmt.Errorf`).
    + if the file is empty, return the error `ErrCorpusIsEmpty`.
    + Use the `strings.Fields` function to get the list of words from the file, and return it.

6. Create a file `corpus_test.go` to test the exposed functionality from the `corpus.go` file. Implement the `TestReadCorpus` function and test different cases (including empty file, etc.)

## v8: picking up a word from the corpus

1. Define the function `pickWord` in `corpus.go` that takes the slice of strings (corpus) and return the selected word as a string.

2. Implement the function as follows:
    + initialize the seed of the random number generator using the current time in nanos.
    + get a random index from the corpus
    + return the word

3. Implement a test for this function. To confirm it works as expected, we will check that a list from the corpus is returned. First of all, define a helper function `inCorpus` that takes the corpus and the word and returns whether it is in there or not.

4. Then, implement the `TestPickWord` in `corpus_internal_test.go` by generating a dummy corpus and validating that we get a random word from the list.

5. Refactor the `New()` function so that it now receives the reader, the corpus and the maxAttempts. (the idea is that the solution will be computed in the implementation from the corpus). Also, the `New` function signature should be enhanced to return an error too.

6. In `New`, start by making sure that the corpus received is not empty, and then, when initializing the game, make sure you're picking a random word from it.

7. Adjust all the files that fail to compile because of the change in `New`.

8. Make the necessary changes in `main.go` so that the corpus is read and passed to the game.

## ToDo


### ByteSlice

Create a type ByteSlice that aliases a `[]byte` and transform the `Append()` function so that it satisfies the `io.Writer` interface.

Confirm that when doing so, you can use methods such as `fmt.Fprintf` to load byteslices.

As an illustration of pointers vs. values, confirm that the compiler allows you to call a pointer method passing a value, but not the other way around.

### Random numbers

Libraries implementing random number generators need to comply with very strict requirements to be considered *safe*. For example, the must generate numbers with the same probability and amount of time.

In Go, the `math/rand` package provides a random number generator, but `crypto/rand` is the one that guarantees truly random numbers, while `math/rand` doesn't. However, the methods in `crypto` are a lot more expensive.

Both packages expose `Intn(n int)` and similar to return random numbers. Also, to truly generate random numbers we'd need to call `Seed(n int)`. By default, the source used to be set as the programmer had used `Seed(1)`. The most common way to get truly random numbers is to pass the time in nanos, although in the newer Go releases the generator is initialized at program startup.

Note that calling `Seed()` more than once in a program is not a good idea.

### Interfaces

A type can implement multiple interfaces. As an example, define a type `Sequence` as an ordered collection of ints (`[]int`).

Then implement the `sort.Interface` for `Sequence`, which requires the definition of `Len`, `Less` and `Swap`.

Define also the `Copy` method that returns a copy of the sequence.

Finally, make Sequence to also implement the *Stringer* interface, and in the implementation, sort the  sequece before printing it.

That is, it should print:

```go
s := Sequence{4, 1, 2, 5, 3}
fmt.Println(s) // [1 2 3 4 5]
```