package functional

import (
	"testing"
)

func TestBind(t *testing.T) {
	var f func() int
	Bind(func(a int) int {
		return a
	}, &f, 1)

	if f() != 1 {
		t.Error("bind error")
	}

	//type args struct {
	//	f    interface{}
	//	args []interface{}
	//}
	//tests := []struct {
	//	name   string
	//	args   args
	//	wantFm interface{}
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if gotFm := Bind(tt.args.f, tt.args.args...); !reflect.DeepEqual(gotFm, tt.wantFm) {
	//			t.Errorf("Bind() = %v, want %v", gotFm, tt.wantFm)
	//		}
	//	})
	//}
}