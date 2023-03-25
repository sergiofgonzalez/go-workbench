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

2. Update the `New()` function so that it receives an `io.Reader` and use it to initialize the `bufio.Reader` of the struct.

3. Define a method `ask()` on `Game` that will read the input from the user until a valid suggestion is made. Once made, it will be returned.

4. Define a constant `wordLength` and make it equal to 5. We will use it to check that the input provided by the user is of the expected length.

5. Implement the `ask()` method. The method should begin by showing a message to to player telling them that it expects a wordLength input. Then, you can implement a while true (for) that will keep asking the user for input until no error is returned by the `Readline()` method. If an error is found, a message should be written in stderr using `Fprintf(os.Stderr)`. Otherwise, if no error is found, you should read the input and place it on a variable `guess` as an array of Unicode points (`[]rune`). HINT: You might need to transform the `[]byte` returned by ReadLine into a string first.

6. In the same method, use `len(guess)` to check that the word typed by the user is of the expected length. Again, if it's not you should display a message of stderr and keep asking the user for input. Otherwise, if it's of the expected length, you should return it.

## v2: testing ask()

1. Create an internal test file in which we will test our `ask()` method (let's call it `game_internal_test.go`).

2. In the file, create a `TestGameAskHappy()` testing method. In it, provide a `strings.NewReader()` when instantiating a `Game` struct. In the asserting, check that the output of `ask()` is the expected array of runes by using the `slices.Equal()` from the slices module (HINT: you will need to do `go get` to add that module).

3. Now, with the `ask()` function tested for the happy path, we can use it in the `Play()` method to get the input from the user. Modidy the method to get the input from the user and then type in stdout the guess.

4. Finally, modify the `main.go` wiring the Stdin to the reader to get the input from the user and check that both the happy path and error path for ask (the loop) work. The program should gracefully terminate when you type a 5 character word, and keep asking for input when typing an invalid input.
