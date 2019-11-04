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
BenchmarkMapRaw-12                            10         166700110 ns/op
BenchmarkMapRawInplace-12                     20          85599920 ns/op
BenchmarkMapper-12                             2         625999000 ns/op
BenchmarkMapperTraits-12                       2         630487700 ns/op
BenchmarkMapperInplaceTraits-12                2         516012900 ns/op
BenchmarkMapperTraits8-12                      5         282200020 ns/op
BenchmarkMapperInplaceTraits4-12              10         146800240 ns/op
BenchmarkMapperInplaceTraits8-12              10         128700010 ns/op
BenchmarkMapRaw4-12                           10         149499990 ns/op
BenchmarkMapRaw8-12                           10         134499610 ns/op

// 100000
BenchmarkMap-12                               50          24939844 ns/op
BenchmarkMapInplace-12                        50          24620058 ns/op


```