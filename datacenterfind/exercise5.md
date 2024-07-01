# Hedged requests

Hedged requests is a concurrency pattern where we want to send a first request, and if the answer is not coming in a given time, then we send the second request, etc.

To do so, we will use a `time.Timer` and play with a `select` statement.

- Create a chan of `bool` for your results.
- Encapsulate the call to `Find` in a goroutine and send the result to the chan.
- Before the goroutine creation, create a `time.Timer` with a duration of `75 milliseconds`.
- After your goroutine, add a `select` statement to select between your result chan and the `timer.C` chan.
- Don't forget to exit your function at the first result.

Test your code:

```bash
go run *go -action=hedged
```

Expected output:

```
2024/06/29 23:43:30 INFO launching find
2024/06/29 23:43:30 INFO launching find
2024/06/29 23:43:30 INFO got result found=true
```

The first Find took too much time, so another Find has been launched. As we finally had a result, there are no other Find launched.