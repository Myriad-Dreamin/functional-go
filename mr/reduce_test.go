package mr

import (
	"fmt"
	"testing"
)

func TestReduce(t *testing.T) {

	var x func(i ...int) []int
	Reduce(func(i int) bool { return i > 2 }, &x)
	var i = []int{4, 7, 1, 2, 9, 1, 10, 7}
	i = x(i...)
	fmt.Print(i)
	//type args struct {
	//	fi interface{}
	//	fm interface{}
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
