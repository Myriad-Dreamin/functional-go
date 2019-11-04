package mr

import (
	"fmt"
	"github.com/Myriad-Dreamin/functional-go"
	"testing"
)

func TestMap(t *testing.T) {
	var x func(i ...int) []int
	Map(func(i int) int {return i+1}, &x)
	var i = []int{1,2}
	i = x(i...)
	fmt.Print(i)
	var y func(i []int) []int
	Map(func(i int) int {return i+1}, &y)
	i = y(i)
	fmt.Print(i)
	//type args struct {
	//	flist interface{}
	//	fi    interface{}
	//	fm    interface{}
	//}
	//tests := []struct {
	//	name string
	//	args args
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//	})
	//}
}

func BenchmarkMap(b *testing.B) {
	var x func(i ...int) []int
	Map(func(i int) int {return i+1}, &x)
	var s = make([]int, 100000)
	for i := 0; i < b.N; i++ {
		s = x(s...)
	}
}


func BenchmarkMapper(b *testing.B) {
	var x func(i ...int) []int
	functional.MakeFunc(Mapper{F: func(sliceI interface{}) Handler {
		slice := sliceI.([]int)
		handler := make(IntHandler, len(slice), cap(slice))
		copy(handler, slice)
		return handler
	},
	}.MapR(func(i int) int {return i+1}), &x)
	var s = make([]int, 100000)
	for i := 0; i < b.N; i++ {
		s = x(s...)
	}
	//print("value:", s[0])
}

func r(i ...int) []int {
	y := make([]int, len(i), cap(i))
	for j := range i {
		y[j] = i[j] + 1
	}
	return y
}

func BenchmarkMapRaw(b *testing.B) {
	var s = make([]int, 100000)
	for i := 0; i < b.N; i++ {
		s = r(s...)
	}
}