package functional

import (
	"reflect"
)

type Function func(in []reflect.Value) (out []reflect.Value)

func BindR(f interface{}, args ...interface{}) Function {
	var fi = reflect.ValueOf(f)
	argsv := Wrap(args...)

	return func(in []reflect.Value) (out []reflect.Value) {
		return fi.Call(append(argsv, in...))
	}
}

func BindRN(f interface{}, n int, args ...interface{}) Function {
	var fi = reflect.ValueOf(f)
	argsv := Wrap(args...)

	return func(in []reflect.Value) (out []reflect.Value) {
		return fi.Call(append(append(in[:n], argsv...), in[n:]...))
	}
}

func Bind(f, fm interface{}, args ...interface{}) {
	var fmi = reflect.ValueOf(fm).Elem()
	fmi.Set(reflect.MakeFunc(fmi.Type(), BindR(f, args...)))
}

func BindN(f, fm interface{}, n int, args ...interface{}) {
	var fmi = reflect.ValueOf(fm).Elem()
	fmi.Set(reflect.MakeFunc(fmi.Type(), BindRN(f, n, args...)))
}


