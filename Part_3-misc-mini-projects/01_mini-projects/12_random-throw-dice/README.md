# Mini-project: Throw the dice
> an exploration of random number generation in Go and using it to simulate the throwing of dice.

## Random numbers

Libraries implementing random number generators need to comply with very strict requirements to be considered *safe*. For example, they must generate numbers with the same probability over the time.

In Go, the `math/rand` package provides a random number generator, but `crypto/rand` is the one that guarantees truly random numbers, while `math/rand` doesn't. However, the methods in `crypto` are a lot heavier.

Both packages expose `Intn(n int)` (and friends) to return random numbers.

They also expose the function `Seed(n int)`, which you can use to drive how the random number generator will behave.

In older versions of Go, the default was to initialize the seed with `Seed(1)`, which is not a very good idea in most cases.

More recent versions automatically seed the random generator at program startup and had deprecated the `Seed` function in favor of a `Seed` method.

The `Seed` method can be used to produce the same output on every run.

##  Project

Create a program that simulates the throwing of a dice and display the results in a frequency table:

```
 2 :  268 (0.027) : **
 3 :  564 (0.056) : ***
 4 :  806 (0.081) : *****
 5 : 1115 (0.112) : ******
 6 : 1432 (0.143) : ********
 7 : 1729 (0.173) : **********
 8 : 1320 (0.132) : ********
 9 : 1076 (0.108) : ******
10 :  800 (0.080) : *****
11 :  573 (0.057) : ***
12 :  317 (0.032) : **
```