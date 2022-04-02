# 目的

本项目测试了4种go字符串拼接方法的性能
1.+ 运算符
2.string.Join
3.bytes.Buffer
4.bytes.Builder

# 运行
go test -bench=. -benchmem

# 结果

```
goos: linux
goarch: amd64
pkg: github.com/chimission/gostringbenchmark
cpu: Intel(R) Core(TM) i5-9500 CPU @ 3.00GHz
BenchmarkStringPlus-3           1000000000               0.2508 ns/op          0 B/op            0 allocs/op
BenchmarkStringJoin-3            8786182               129.3 ns/op            32 B/op            1 allocs/op
BenchmarkStringBuffer-3          8927844               143.2 ns/op            96 B/op          2 allocs/op
BenchmarkStringBuilder-3         8316704               133.3 ns/op            56 B/op          3 allocs/op
BenchmarkStringPlus2-3             25015             53991 ns/op          161840 B/op         99 allocs/op
BenchmarkStringJoin2-3            555967              2230 ns/op            3072 B/op          1 allocs/op
BenchmarkStringBuffer2-3          247363              5442 ns/op           15168 B/op          8 allocs/op
BenchmarkStringBuilder2-3         328734              4090 ns/op           10464 B/op         10 allocs/op
PASS
ok      github.com/chimission/gostringbenchmark 14.161s
```