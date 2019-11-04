package mr

import (
	"github.com/Myriad-Dreamin/functional-go"
	"reflect"
)

type MapperF interface {
	F(slice interface{}) Handler
	FInplace(slice interface{}) Handler
}

type Mapper struct {
	MapperF
	CoreCount int
}

func (m Mapper) MapR(fi interface{}) functional.Function {
	if m.CoreCount == 0 {
		m.CoreCount = 1
	}
	t := reflect.SliceOf(reflect.TypeOf(fi).In(0))
	return func(in []reflect.Value) []reflect.Value {
		slice := in[0]
		in[0] = reflect.ValueOf(MapSlice(m.F(slice.Interface()),
			0, slice.Len(), m.CoreCount, fi)).Convert(t)
		return in
	}
}

func (m Mapper) MapRInplace(fi interface{}) functional.Function {
	if m.CoreCount == 0 {
		m.CoreCount = 1
	}
	t := reflect.SliceOf(reflect.TypeOf(fi).In(0))
	return func(in []reflect.Value) []reflect.Value {
		slice := in[0]
		in[0] = reflect.ValueOf(MapSlice(m.FInplace(slice.Interface()),
			0, slice.Len(), m.CoreCount, fi)).Convert(t)
		return in
	}
}
