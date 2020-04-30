# FastestSumOfSquares

## Running

1. Download the repo.
2. Run using `go run sumOfSquares.go`

## Proof that concurrent solution works faster!
```
solve took 322ns
solve took 99.558µs
solve took 226ns
concurrent_solve took 45.729µs
5
5
```
## References

1. https://medium.com/@_orcaman/when-too-much-concurrency-slows-you-down-golang-9c144ca305a
2. https://medium.com/@Mardiniii/make-it-real-elite-week-2-recursivity-concurrency-and-goroutines-e740bd4ca780
3. https://stackoverflow.com/questions/35588474/what-does-allocs-op-and-b-op-mean-in-go-benchmark
4. https://coderwall.com/p/cp5fya/measuring-execution-time-in-go
5. https://blog.learngoprogramming.com/5-gotchas-of-defer-in-go-golang-part-ii-cc550f6ad9aa
6. https://www.youtube.com/watch?v=YEKjSzIwAdA
