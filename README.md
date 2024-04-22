# Bubble Bench

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