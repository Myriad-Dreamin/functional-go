package mr

import (
	"fmt"
	"github.com/Myriad-Dreamin/functional-go"
	"testing"
)

const factor = 100000000

func add1(i int) int { return i + 1 }

func TestMap(t *testing.T) {
	var x, y, z func(i ...int) []int
	var i = []int{1, 2}

	Map(func(i int) int { return i + 1 }, &x)
	//i = x(i...)
	//fmt.Println(i)

	Map(add1, &y)
	//i = y(i...)
	//fmt.Println(i)

	var mapper = NewMapperTraits(IntHandler{})
	mapper.Map(add1, &z)
	i = z(i...)
	fmt.Println(i)

	var mapper2 = NewMapperTraits(IntHandler{}, 2)
	mapper2.Map(add1, &z)
	i = z(i...)
	fmt.Println(i)

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

func BenchmarkMapRaw(b *testing.B) {
	var s = make([]int, factor)
	for i := 0; i < b.N; i++ {
		s = mapF(add1, s...)
	}
	//fmt.Println("|", s[0], "|")
}

func BenchmarkMapRawInplace(b *testing.B) {
	var s = make([]int, factor)
	for i := 0; i < b.N; i++ {
		s = mapFInplace(add1, s...)
	}
	//fmt.Println("|", s[0], "|")
}

func BenchmarkMapper(b *testing.B) {
	var x func(i ...int) []int
	functional.MakeFunc(IntMapper{}.MapR(func(i int) int { return i + 1 }), &x)
	var s = make([]int, factor)
	for i := 0; i < b.N; i++ {
		s = x(s...)
	}
	//fmt.Println("|", s[0], "|")
}

func BenchmarkMapperTraits(b *testing.B) {
	var x func(i ...int) []int
	functional.MakeFunc(NewMapperTraits(IntHandler{}).MapR(func(i int) int { return i + 1 }), &x)
	var s = make([]int, factor)
	for i := 0; i < b.N; i++ {
		s = x(s...)
	}
	//fmt.Println("|", s[0], "|")
}

func BenchmarkMapperInplaceTraits(b *testing.B) {
	var x func(i ...int) []int
	functional.MakeFunc(NewMapperTraits(IntHandler{}).MapRInplace(func(i int) int { return i + 1 }), &x)
	var s = make([]int, factor)
	for i := 0; i < b.N; i++ {
		s = x(s...)
	}
	//fmt.Println("|", s[0], "|")
}

func BenchmarkMapperTraits8(b *testing.B) {
	var x func(i ...int) []int
	functional.MakeFunc(NewMapperTraits(IntHandler{}, 8).MapR(func(i int) int { return i + 1 }), &x)
	var s = make([]int, factor)
	for i := 0; i < b.N; i++ {
		s = x(s...)
	}
	//fmt.Println("|", s[0], "|")
}

func BenchmarkMapperInplaceTraits4(b *testing.B) {
	var x func(i ...int) []int
	functional.MakeFunc(NewMapperTraits(IntHandler{}, 4).MapRInplace(func(i int) int { return i + 1 }), &x)
	var s = make([]int, factor)
	for i := 0; i < b.N; i++ {
		s = x(s...)
	}
	//fmt.Println("|", s[0], "|")
}

func BenchmarkMapperInplaceTraits8(b *testing.B) {
	var x func(i ...int) []int
	functional.MakeFunc(NewMapperTraits(IntHandler{}, 8).MapRInplace(func(i int) int { return i + 1 }), &x)
	var s = make([]int, factor)
	for i := 0; i < b.N; i++ {
		s = x(s...)
	}
	//fmt.Println("|", s[0], "|")
}

func BenchmarkMapRaw4(b *testing.B) {
	var s = make([]int, factor)
	for i := 0; i < b.N; i++ {
		s = MapSlice(IntHandler(s), 0, factor, 4, func(a int) int { return a + 1 }).(IntHandler)
	}
	//fmt.Println("|", s[0], "|")
}

func BenchmarkMapRaw8(b *testing.B) {
	var s = make([]int, factor)
	for i := 0; i < b.N; i++ {
		s = MapSlice(IntHandler(s), 0, factor, 8, func(a int) int { return a + 1 }).(IntHandler)
	}
	//fmt.Println("|", s[0], "|")
}

func mapF(f IntMapFunc, i ...int) []int {
	y := make([]int, len(i), cap(i))
	for j := range i {
		y[j] = f(i[j])
	}
	return y
}

func mapFInplace(f IntMapFunc, i ...int) []int {
	for j := range i {
		i[j] = f(i[j])
	}
	return i
}
