# Bubble Bench

> Running on a **64-bit** **Linux** machine with **13th Gen Intel(R) Core(TM) i5-13500H** CPU

- _N_: Length of array to be sorted.
- _Growth_: Duration of current row devided by duration of previous row.

## Bubble Sort

| N       | Operations | Duration           | Growth    |
| ------- | ---------- | ------------------ | --------- |
| 100     | 245686     | 4435 ns/op         | -         |
| 1000    | 5431       | 221913 ns/op       | 5003.68%  |
| 10000   | 57         | 20505701 ns/op     | 9240.42%  |
| 100000  | 1          | 2076236080 ns/op   | 10125.17% |
| 1000000 | 1          | 217910336838 ns/op | 10495.45% |

## Fake Sort

| N       | Operations | Duration           | Growth    |
| ------- | ---------- | ------------------ | --------- |
| 100     | 272680     | 4479 ns/op         | -         |
| 1000    | 5535       | 223233 ns/op       | 4983.99%  |
| 10000   | 57         | 20534557 ns/op     | 9198.71%  |
| 100000  | 1          | 2074326401 ns/op   | 10101.64% |
| 1000000 | 1          | 219491639306 ns/op | 10581.35% |

## Raw Results

Here is the raw output of `go test ./... -bench=.` command

```
goos: linux
goarch: amd64
pkg: github.com/utilyre/bubblebench/pkg/sort
cpu: 13th Gen Intel(R) Core(TM) i5-13500H
BenchmarkBubbleSort_n100-16               245686              4435 ns/op
BenchmarkBubbleSort_n1000-16                5431            221913 ns/op
BenchmarkBubbleSort_n10000-16                 57          20505701 ns/op
BenchmarkBubbleSort_n100000-16                 1        2076236080 ns/op
BenchmarkBubbleSort_n1000000-16                1        217910336838 ns/op
BenchmarkFakeSort_n100-16                 272680              4479 ns/op
BenchmarkFakeSort_n1000-16                  5535            223233 ns/op
BenchmarkFakeSort_n10000-16                   57          20534557 ns/op
BenchmarkFakeSort_n100000-16                   1        2074326401 ns/op
BenchmarkFakeSort_n1000000-16                  1        219491639306 ns/op
```
