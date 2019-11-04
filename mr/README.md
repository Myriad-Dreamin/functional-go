# Map Reduce Function Tools

## Benchmark
```plain
GOROOT=/home/kamiyoru/work/go #gosetup
GOPATH=/home/kamiyoru/work/gosrc #gosetup
/home/kamiyoru/work/go/bin/go test -c -o /tmp/___gobench_github_com_Myriad_Dreamin_functional_go_mr github.com/Myriad-Dreamin/functional-go/mr #gosetup
/tmp/___gobench_github_com_Myriad_Dreamin_functional_go_mr -test.v -test.bench . -test.run ^$ #gosetup
goos: linux
goarch: amd64
pkg: github.com/Myriad-Dreamin/functional-go/mr
BenchmarkMap-4                    	      30	  36424100 ns/op
BenchmarkMapInplace-4             	      30	  35343890 ns/op
BenchmarkMapRaw-4                 	    5000	    330581 ns/op
BenchmarkMapRawInplace-4          	   20000	     64204 ns/op
BenchmarkMapper-4                 	    1000	   1239542 ns/op
BenchmarkMapperTraits-4           	    1000	   1410961 ns/op
BenchmarkMapperInplaceTraits-4    	    2000	    630673 ns/op
BenchmarkMapperTraits8-4          	    5000	    378276 ns/op
BenchmarkMapperInplaceTraits4-4   	  200000	      7324 ns/op
BenchmarkMapperInplaceTraits8-4   	  200000	      9971 ns/op
BenchmarkMapRaw4-4                	  500000	      3685 ns/op
PASS

Process finished with exit code 0
```