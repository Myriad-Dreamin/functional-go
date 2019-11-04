package mr

import "testing"

func BenchmarkMap(b *testing.B) {
	var x func(i ...int) []int
	Map(func(i int) int {return i+1}, &x)
	var s = make([]int, 100000)
	for i := 0; i < b.N; i++ {
		s = x(s...)
	}
}

func BenchmarkMapInplace(b *testing.B) {
	var x func(i ...int) []int
	Map(func(i int) int {return i+1}, &x)
	var s = make([]int, 100000)
	for i := 0; i < b.N; i++ {
		x(s...)
	}
}

