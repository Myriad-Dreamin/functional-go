package mr

import (
	"github.com/Myriad-Dreamin/functional-go"
	"reflect"
	"sync"
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

func MapRInplace(fi interface{}) functional.Function {
	var f = reflect.ValueOf(fi)
	return func(in []reflect.Value) []reflect.Value {
		slice := in[0]
		for i := slice.Len()-1; i >=0;i--  {
			in[0] = slice.Index(i)
			slice.Index(i).Set(f.Call(in)[0])
		}
		in[0] = slice
		return in
	}
}

func Map(fi, fm interface{})  {
	functional.MakeFunc(MapR(fi), fm)
	return
}

func MapInplace(fi, fm interface{})  {
	functional.MakeFunc(MapRInplace(fi), fm)
	return
}

type Handler interface {
	Call(f interface{}, index int)
	Value() reflect.Value
}

type MapperF interface {
	F(slice interface{}) Handler
	FInplace(slice interface{}) Handler
}

type Mapper struct {
	MapperF
	CoreCount int
}

func mapSlice(handler Handler, l, r int, f interface{}) {
	for i := l; i < r;i ++  {
		handler.Call(f, i)
	}
}

func min(l, r int) int {
	if l < r {
		return l
	}
	return r
}

func MapSlice(handler Handler, l, r, coreCount int, f interface{}) Handler {
	step := (r-l+coreCount-1)/coreCount
	if coreCount <= 1 || step <= 0 {
		mapSlice(handler, l, r, f)
	} else {
		var wg sync.WaitGroup
		wg.Add(coreCount)
		for i := l; i < r; i += step {
			go func(i int) {
				mapSlice(handler, i, min(i + step, r), f)
				wg.Done()
			}(i)
		}
		wg.Wait()
	}
	return handler
}

func (m Mapper) MapR(fi interface{}) functional.Function {
	if m.CoreCount == 0 {
		m.CoreCount = 1
	}
	return func(in []reflect.Value) []reflect.Value {
		slice := in[0]
		in[0] = MapSlice(m.F(slice.Interface()),
			0, slice.Len(), m.CoreCount, fi).Value()
		return in
	}
}

func (m Mapper) MapRInplace(fi interface{}) functional.Function {
	if m.CoreCount == 0 {
		m.CoreCount = 1
	}
	return func(in []reflect.Value) []reflect.Value {
		slice := in[0]
		in[0] = MapSlice(m.FInplace(slice.Interface()),
			0, slice.Len(), m.CoreCount, fi).Value()
		return in
	}
}

