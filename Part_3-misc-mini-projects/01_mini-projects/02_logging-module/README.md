# Mini-project: a custom logging module

## v0

1. Start by creating the project's dir that will host our logging source code. In Go, that is done by creating a directory with the name of the project (e.g., `logging-module`)-

2. Initialize the Go module by running `go mod init`. Use the name `example.com/pocketlog` for the module name.

3. Create a directory `pocketlog/` to host the package.

4. Create a file `logger.go` in the `pocketlog/` dir.

5. Define a struct `Logger` that will contain the fields of the logger.

6. Define another file `level.go` in the same package. Declare a type `Level` that aliases the `byte` type. We won't need to support a lot of levels, and therefore `byte` will use the right amount of memory.

7. In the same `level.go` file, create an enumeration for the logging levels with names `LevelDebug`, `LevelInfo`, `LevelError`. Do this by creating const variables, and assigning the value `= iota` to the first value in the enumeration.

| NOTE: |
| :---- |
| Using `= iota` lets the compiler know that we are starting an enumeration. It allows us to create a sequence of numbers incremented on each line, without having to assign explicit values to these constants. Iota can be used on any type that is based on an integer. The values will be automatically incremented on every line, so the values should be sorted according to this convention. |

8. In `logger.go`, define the methods to accomodate the APIs: `l.Logf(pocketlog.Info, "message")`, `l.Debugf("message")`, `l.Infof("message")`, `l.Errorf("message")`.

9. Then define also the methods that mimic how `Printf` works by defining variations such as `l.Debugf(format string, args ...any)`.

10. Create a `main.go` to validate the functionality created this far. No tests needed for now.

## v1

Go does not provide support for constructors, but it's a good practive to define a `New()` method that returns a new instances and simplifies the cognitive load on the user as they wouldn't need to do things such as:

```go
var l1 pocketLog.Logger

l2 := pocketLog.Logger{}
```

Also: We have created the `Level` type, but we're still haven't defined the corresponding logging threshold level that supresses all the log lines below the corresponding threshold:

```
LevelDebug < LevelInfo < LevelError
```

The idea is that if we set the level to `LevelDebug`, all logging information will be displayed, if we set `LevelInfo`, debug log lines will be suppressed, and if we set `LevelError`, only error lines will be displayed.

1. Define a field in our `Logger` struct called `threshold` to keep the level of the logger.

2. Define a function `New()` that accepts the threshold, and returns a pointer to an initialized Logger value.

| NOTE: |
| :---- |
| We made the decision to return a pointer to a Logger rather than a Logger instance. We are mimicking what C++ does with its built-in operator, and it's generally more useful. |

3. Adjust the existing implementations to be aware of the threshold and suppress loglines accordingly.

4. Test informally in `main()` using `New()` and different thresholds.
Note that the user can still use the struct to initialize the Logger. Investigate what is the default value of the threshold when the user do not use `New()` to initialize it.

## v2

In this iteration we will be adding proper tests. As we're building a module that will be potentially consumed by other users we should be concerned with *closed-box testing*, that is, testing a system from the outside, just as a user would interact with our module.

1. Start by creating a `logger_test.go`. Place it in the same `pocketlog/` directory, but use a different package naming because it shouldn't be part of the module.

    This ensures that the test code that we put in `logger_test.go` will only have access to the `pocketlog`'s package public interface.

2. Use the `Example_*` approach to test the `Logf`, `Debugf`, `Infof`, and `Errorf` methods.

3. Clean `main.go` as we now have proper tests

## v3

There is an unofficial convention to write a special file in each Go package that will describe the purpose of this package. This file is named `doc.go`, and is called a package header.

The `doc.go` must contain no Go code, except for one uncommented line: the package. Before that line, you should include a verbose description of what the package is about.


1. Create a `go doc` file.

2. Give a try to the `go doc` command. This command will give you the documentation of a package or a symbol that the `go` comand can find in the subdirectories. Run `go doc path/to/pocketlog/` to see the package documentation, and `go doc path/to/pocketlog.New` to get information about the `New` function and `go doc path/to/pocketlog.Logger.Logf` to get info about `Logf`. Can you spot how `go doc` uses `Example_` test functions?

| NOTE: |
| :---- |
| You will need to use the absolute path. |

## v4

Enhance the logging framework so that the Logger can be created with a custom variable of type `io.Writer`. The idea is to allow the end user to define the destination of their choice.

The `io.Writer` interface allows you to write to any destination and it is declared as follows:

```go
type Writer interface {
    Write(p [byte]) (n int, err error)
}
```

1. Make a modification of the `Logger` definition to also hold the selected destination (which should be of type `io.Writer`).

2. Update the implementation of the methods so that it uses the new Logger field. When the value of the destination is `nil`, assign the standard output to it using `os.Stdout`. You will need to use `FprintF()` which requires a variable of type `io.Writer`. Make sure you use an explicit approach for getting the values returned by `FprintF()` using `_`.

3. Fix the tests to accommodate the changes.

## v5

While we want to reduce the cognitive load when using our logger, we also want to support an extensible logger that can support multiple destinations (and other potential features such as datatime formatting).

One approach to support that is by using *configuration functions*.

1. Create an `options.go` in the `pocketlog` package. In it, define the type `Option` representing functions that take a `*Logger` as parameter.

2. In the same file, define the function `WithOutput()` that takes the desired output and returns an `Option`. Implement it so that the end user can provide the desired destination in an easy way.

3. Modify the `New()` function so that it accepts a *variadic* `options` parameter. In the implementation, iterate over those options calling them to configure the logger.

4. Update `main.go` and the tests to accommodate the changes.


## v6

We have fixed the tests and validated that it works well for the stdout, but we have never tested with other outputs.

As we are requiring an interface of type `io.Writer` as argument, nothing prevents us from creating a test implementation of such an interface so that we can fully test that capability.

1. In the existing `logger_test.go` file, create a `TestWriter` struct with a single field `contents` of type string.

2. Make `TestWriter` implement `io.Writer` interface.

3. Create specific tests that no longer rely on the `Example_` approach, but rather use a Logger that has the `TestWriter` as the output. Create individual tests to assert the functionality of the logger (NOTE: those will be generalized in the next version).

## v7

With the recent test approach there's a lot of repetition, and most of the tests look the same.

1. Refactor the existing tests using `t.Run()`. You can keep an `Example_` test just to validate the basics, but you should use the approach introduced in [v6](#v6) to get good coverage of the Logger functionality.

2. Check the code coverage of your tests by running: `go test -coverprofile=coverage.out` followed by `go tool cover -html=c.out`

## v8

Let's enhance our logger with a few best practices with respect to logging:

1. Enhance the logger to ensure that each log line does not exceed 1000 characters (or bytes), and allow the user to fine-tune this limit.

2. Enhance the logs so that it can optionally display the time at which the log line was emitted. By default, no time information will be used.

3. Enhance the logger so that it can produce logs as events in JSON with the following structure:
  + `level` &mdash; numeric value for the log level
  + `time` &mdash; the time at which the log was emitted using ticks.
  + `msg` &mdash; the message with all the values embedded into the string.
  
4. Update the tests to cover this new functionality.