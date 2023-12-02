# Advent 2023

# Benchmarks 2023
Run the following to show benchmarks on your machine.
```sh
go test ./... -bench=.
```

Explanation: Day 1, part 1 benchmark was run 31356 times, averaging 38.5us runtime. Per run 48 bytes were allocated in a single allocation.
The benchmark is run against  the actual input.
```
goos: linux
goarch: amd64
pkg: github.com/soypat/advent/2023/1
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz
BenchmarkPartOne-16        32408             37650 ns/op              48 B/op          1 allocs/op
BenchmarkPartTwo-16        26841             44594 ns/op              48 B/op          1 allocs/op
PASS
ok      github.com/soypat/advent/2023/1 3.248s
elapsed: 19.165µs
elapsed: 56.387µs
elapsed: 9.083µs
elapsed: 37.067µs
goos: linux
goarch: amd64
pkg: github.com/soypat/advent/2023/2
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz
BenchmarkPartOne-16        54613             23131 ns/op              48 B/op          1 allocs/op
BenchmarkPartTwo-16        39333             30090 ns/op              48 B/op          1 allocs/op
PASS
ok      github.com/soypat/advent/2023/2 2.981s
```