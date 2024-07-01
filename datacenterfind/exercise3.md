# Redundant

Let's now imagine another use case: you just need the first result of all your calls, wherever it comes from.

To do so, we will use the redundant pattern: starting all the calls concurrently, wait for the first result and cancel all other running jobs.

The [context.Context structure](https://pkg.go.dev/context) is designed for this purpose, thanks to the [WithCancel](https://pkg.go.dev/context#WithCancel) mathod.

To understand how cancelation is handled, you can have a look on the `FindWithContext` method definition in [finder.go](finder.go).

Open the [redundant.go](./redundant.go) file:

- Create a new cancelable context from `context.Background()`
- Fix the `context.TODO()` with this new context
- Stop the function after the first result, then call the cancel method

Test your code:

```bash
go run *go -action=redundant
```

Expected output:

```
2024/06/29 23:05:07 INFO deadline exceeded finder=GRA
2024/06/29 23:05:07 INFO deadline exceeded finder=BHS
2024/06/29 23:05:07 INFO got result datacenter=SBG found=true
```

You should got result from one datacenter, and two others with a deadline exceeded log.