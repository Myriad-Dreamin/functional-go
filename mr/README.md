# Map Reduce Function Tools

## Benchmark
```plain
goos: linux
goarch: amd64
pkg: github.com/Myriad-Dreamin/functional-go/mr
// 100000000
BenchmarkMapRaw-4                 	       2	 701370729 ns/op
BenchmarkMapRawInplace-4          	       5	 210124306 ns/op
BenchmarkMapper-4                 	       1	1199713562 ns/op
BenchmarkMapperTraits-4           	       1	1189000453 ns/op
BenchmarkMapperInplaceTraits-4    	       2	 728959362 ns/op

// 100000
BenchmarkMap-4                    	      30	  35993888 ns/op
BenchmarkMapInplace-4             	      50	  34801953 ns/op

```