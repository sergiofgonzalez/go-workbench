# Mocking in Go using TDD
> this example illustrates how to use mocking in Go


You have been asked to write a program which counts down from 3, printing each number on a new line, with a 1-second pause between each lines, and when it reaches 0 it will print "Go!"

```
3
2
1
Go!
```

This behavior will be encoded in `main`, so that you get that behavior when you run the program, and get the output on stdout.


## v0: setting the stage

By looking at the requirements it is clear that we need to end up having something like:

```go
func main() {
  Countdown()
}
```

We will approach the implementation of the `Countdown` in several stages.

1. Make `Countdown` just display 3
2. Make `Countdown` display 3, 2, 1, Go!
3. Make `Countdown` to wait for 1 second between each line

So in this version, we will start with #1.

## v1: Getting the whole countdown

Update the previous version, so that now `Countdown` displays the whole count, that is, implement #2.

## v2: Introducing the sleep

In this version we can accommodate the `sleep` associated with #3. Besides, it will be a good time to refactor a little bit the implementation of `Countdown` removing the magic numbers, etc.

## v3: Introducing mocking

At this point we have a working, testable version but there is something specific we need to improve:

> the tests take some time to run, because it requires at least 3 seconds. What if the sleep were 1 minute or one hour?

In order to fix that, we need to *extract* the dependency on `Sleep`, apply *DI*, and in our tests use provide a *mock* instance instead of the real `time.Sleep()`, which we will use in outside of tests.

You typically will want to start defining an interface:

```go
type Sleeper interface {
  Sleep()
}
```

Note that in this first approach, I'm not interested in how long Sleep will delay the execution.

Instead, I will focus in how many times `Sleep` is called. For that I neeed to use a spy:

```go
type SpySleeper struct {
  Calls int
}

func (s *SpySleeper) Sleep() {
  s.Calls++
}
```

With all this information in place, we can make the necessary changes in the project.

A few hints:

Start by defining the `Sleeper` interface in the main program. Along with that, to be able to invoke the real `time.Sleep` you will need to create a struct that implements `Sleeper` and that can be forwarded to `Countdown` to sleep for one second between the different displays.

Then, define the `SpySleeper` in your test code, and make sure you assert that the number of calls match your expectations.

## v4: Improving your tests

The previous version is running as expected, but we're not confident that our tests really ensure the functionality we're expecting.

To confirm it, start by making the following implementation change in `Countdown` and confirm that the tests pass:

```go
func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
	}

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
	}

	fmt.Fprintln(w, "Go!")
}
```

That is not the sequence we're expecting, and yet, when we run the tests we get:

```bash
$ go test
PASS
ok      example.com/10_tdd-mocking      0.001s
```

This is an indication that tests should be improved.

This can be fixed by improving our spy to not only track the number of calls, but also the sequence.

Let's fix it by creating another spy that records the sequence of calls, so that when we run our test we get:

```
--- FAIL: TestCountdown (0.00s)
    --- FAIL: TestCountdown/confirm_test_sequence (0.00s)
        countdown_test.go:58: got [sleep sleep sleep write write write write], but wanted [write sleep write sleep write sleep write]
FAIL
exit status 1
FAIL    example.com/10_tdd-mocking      0.001s
```

To make it work you just need to create another *spy* `SpyCountdown` that can be injected into `Countdown` to keep track of the sequence of calls.

That is really simple to accomplish once you know how:

```go
type SpyCountdown struct {
	Calls []string
}

func (s *SpyCountdown) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func (s *SpyCountdown) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return
}
```

Then, you can use it in your tests:

```go
t.Run("confirm test sequence", func(t *testing.T) {
  spyCountdown := SpyCountdown{}
  want := []string{"write", "sleep", "write", "sleep", "write", "sleep", "write"}

  Countdown(&spyCountdown, &spyCountdown)
  if !slices.Equal(want, spyCountdown.Calls) {
    t.Errorf("got %v, but wanted %v", spyCountdown.Calls, want)
  }
})
```

See how the same instance is wired for both the writer and the sleeper, and that is why we created both the `Sleep` and `Write` implementations for our spy.

## v5: adjusting the sleep duration

Right now the application is working as per requirements, and the only thing we're not checking is that the sleep duration is 1-second.

To finalize the project, we will do so.

We will be substituting the existing `DefaultSleeper` that we injected into `Countdown` for a more flexible `ConfigurableTimeSleeper`:

1. Create a `ConfigurableTimeSleeper` struct that has a `duration` and `sleepFn` fields.

2. Make `ConfigurableTimeSleeper` implement the `Sleeper` interface. That way, we don't need to change the `Countdown` function.

3. In `main`, create composite literal of type `ConfigurableTimeSleeper` with a duration of 1 second, and configured to use `time.Sleep` as the sleep function.

4. Then, write tests for the `ConfigurableTimeSleeper`. This will require you to create another spy that checks that it sleeps for the configured time.