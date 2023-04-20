# Hello, World TDD
> this example illustrates how to set up shop with TDD in Go, in a quite detailed manner so that we don't need to get into such details for the subsequent ones.

The TDD loop is:
1. Write a test

2. Write code in your program to make the compiler pass

3. Make it fail, and confirm that the error messages help identify the problem.

4. Write enough code to make the test pass

5. Refactor

## v0: Write a test

In this version you have a module with an empty `main` package in which you create a `main_test.go`.

Within the test you create a test for the output of `main()` as described in the requirements, and another test for the `Hello` function.

The code shouldn't compile as there's no `Hello` function for now.

When running the test it should be clear that there is code to be included.

```bash
$ go test .
# example.com/helloworld [example.com/helloworld.test]
./main_test.go:14:10: undefined: Hello
./main_test.go:20:10: undefined: Hello
./main_test.go:26:10: undefined: Hello
FAIL    example.com/helloworld [build failed]
FAIL
```

## v1: make the compiler pass + confirm error messages

In this version you just need to write enough code to make the test compile.

When doing `go test .` it should be clear that all tests are failing and the errors should be clear enough:

```
$ go test .
--- FAIL: TestHello (0.00s)
    --- FAIL: TestHello/no_name_given (0.00s)
        main_test.go:15: got "", but wanted "Hello, World!"
    --- FAIL: TestHello/'Jason_Isaacs'_name_given (0.00s)
        main_test.go:21: got "", but wanted "Hello to Florence Pugh!"
    --- FAIL: TestHello/'Florence_Pugh'_name_given (0.00s)
        main_test.go:27: got "", but wanted "Hello to Florence Pugh!"
--- FAIL: Example_main (0.00s)
got:

want:
Hello to Jason Isaacs!
FAIL
FAIL    example.com/helloworld  0.001s
FAIL
```

## v2: write enough code to make the tests pass

In this version you just need to write enough code to make the test compile.

## v3: refactor

In this version, you confirm what can be done in the tests, and the program in order to make the whole project more readable and maintainable.

For example, we can create a test structure to formalize how we approach the testing of the `Hello` function to make it more extensible going forwards.