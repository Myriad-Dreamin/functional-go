package mr

import (
	"github.com/Myriad-Dreamin/functional-go"
	"reflect"
)

func MapR(fi interface{}) functional.Function {
	var f = reflect.ValueOf(fi)
	return func(in []reflect.Value) (out []reflect.Value) {
		slice := in[0]
		sliceMap := reflect.MakeSlice(slice.Type(), slice.Len(), slice.Cap())
		for i := slice.Len() - 1; i >= 0; i-- {
			in[0] = slice.Index(i)
			sliceMap.Index(i).Set(f.Call(in)[0])
		}
		out = append(out, sliceMap)
		return out
	}
}

func MapRInplace(fi interface{}) functional.Function {
	var f = reflect.ValueOf(fi)
	return func(in []reflect.Value) []reflect.Value {
		slice := in[0]
		for i := slice.Len() - 1; i >= 0; i-- {
			in[0] = slice.Index(i)
			slice.Index(i).Set(f.Call(in)[0])
		}
		in[0] = slice
		return in
	}
}

func Map(fi, fm interface{}) {
	functional.MakeFunc(MapR(fi), fm)
	return
}

func MapInplace(fi, fm interface{}) {
	functional.MakeFunc(MapRInplace(fi), fm)
	return
}
