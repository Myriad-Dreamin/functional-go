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


goos: windows
goarch: amd64
pkg: github.com/Myriad-Dreamin/functional-go/mr
// 100000000
BenchmarkMapRaw-12                  	       3	 371668433 ns/op
BenchmarkMapRawInplace-12           	       5	 240799760 ns/op
BenchmarkMapper-12                  	       2	 628998950 ns/op
BenchmarkMapperTraits-12            	       2	 623999250 ns/op
BenchmarkMapperInplaceTraits-12     	       2	 513511500 ns/op
BenchmarkMapperTraits8-12           	       5	 281000660 ns/op
BenchmarkMapperInplaceTraits4-12    	      10	 158599680 ns/op
BenchmarkMapperInplaceTraits8-12    	      10	 126300010 ns/op
BenchmarkMapRaw4-12                 	      10	 147497270 ns/op
BenchmarkMapRaw8-12                 	      10	 126496710 ns/op

// 100000
BenchmarkMap-12                               50          24939844 ns/op
BenchmarkMapInplace-12                        50          24620058 ns/op


```