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
		for i := slice.Len()-1; i >=0;i--  {
			in[0] = slice.Index(i)
			sliceMap.Index(i).Set(f.Call(in)[0])
		}
		out = append(out, sliceMap)
		return out
	}
}

func Map(fi, fm interface{})  {
	var fmi = reflect.ValueOf(fm).Elem()
	fmi.Set(reflect.MakeFunc(fmi.Type(), MapR(fi)))
}



type Handler interface {
	Call(f interface{}, index int)
	Value() reflect.Value
}

type Mapper struct {
	F func(slice interface{}) Handler
}

type IntHandler []int

type IntMapFunc = func(int) int

func (handler IntHandler) Call(f interface{}, index int) {
	handler[index] = f.(IntMapFunc)(handler[index])
}

func (handler IntHandler) Value() reflect.Value {
	return reflect.ValueOf(([]int)(handler))
}

func (m Mapper) MapR(fi interface{}) functional.Function {
	return func(in []reflect.Value) (out []reflect.Value) {
		slice := in[0]
		sliceMap := m.F(slice.Interface())
		for i := slice.Len()-1; i >=0;i--  {
			sliceMap.Call(fi, i)
		}
		out = append(out, sliceMap.Value())
		return out
	}
}

