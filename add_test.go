package functional

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var f func(a, b int) int
	Add(&f)
	fmt.Println(f(1,2))
	AddC(&f)
	fmt.Println(f(1,2))
	//type args struct {
	//	f interface{}
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

func BenchmarkAdd(b *testing.B) {
	var f func(a, b int) int
	Add(&f)
	for i := 0; i < b.N; i++ {
		_ = f(1,1)
	}
}