package mr

import (
	"github.com/Myriad-Dreamin/functional-go"
	"reflect"
)

func ReduceR(fi interface{}) functional.Function {
	var f = reflect.ValueOf(fi)
	return func(in []reflect.Value) (out []reflect.Value) {
		slice := in[0]
		sliceReduce := reflect.MakeSlice(slice.Type(), slice.Len(), slice.Cap())
		l,r := slice.Len(), 0
		for i := 0; i < l; i++  {
			in[0] = slice.Index(i)
			if f.Call(in)[0].Bool() {
				sliceReduce.Index(r).Set(in[0])
				r++
			}
		}
		out = append(out, sliceReduce.Slice(0, r))
		return out
	}
}

func Reduce(fi, fm interface{})  {
	var fmi = reflect.ValueOf(fm).Elem()
	fmi.Set(reflect.MakeFunc(fmi.Type(), ReduceR(fi)))
}

