# Map Reduce Function Tools

## Benchmark
```plain
goos: linux
goarch: amd64
pkg: github.com/Myriad-Dreamin/functional-go/mr
BenchmarkMap-4                    	      30	  35166203 ns/op
BenchmarkMapInplace-4             	      50	  35727018 ns/op
BenchmarkMapRaw-4                 	       2	 604567604 ns/op
BenchmarkMapRawInplace-4          	       5	 205347513 ns/op
BenchmarkMapper-4                 	       1	1183385092 ns/op
BenchmarkMapperTraits-4           	       1	1189837245 ns/op
BenchmarkMapperInplaceTraits-4    	       2	 725750223 ns/op
BenchmarkMapperTraits8-4          	       2	 895631293 ns/op
BenchmarkMapperInplaceTraits4-4   	       3	 403120648 ns/op
BenchmarkMapperInplaceTraits8-4   	       3	 431444227 ns/op
BenchmarkMapRaw4-4                	   20000	     80905 ns/op

```