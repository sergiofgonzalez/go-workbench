# Mini-project: A minimalist tracing library using defer

The goal of this mini-project is to create a library `tracing` with a couple of functions `Trace` and `Un` that can be used to trace the execution flow of a program.

The DX is based on using defer to trigger the tracing enablement and deactivation:

```go
func f() {
  defer Tracing.Un(Tracing.Trace("f"))
  fmt.Println("...doing important stuff in f...")
}
```

As the argument to the deferred function `Tracing.Un` are computed when `defer` is being executed, a call to `Tracing.Trace` will occur at the beginning of the function. Then, the deactivation will happen right after exiting `f` thus resulting in the output:

```
In f
...doing important stuff in f....
Exiting f
```

This approach could be enhanced with timestamps to measure execution time, etc.

The implementation is based on an example from [Effective Go](https://go.dev/doc/effective_go#defer)