package functional

import "reflect"

func Wrap(args ...interface{}) []reflect.Value {
	var argsv = make([]reflect.Value, len(args))
	for i := range args {
		argsv[i] = reflect.ValueOf(args[i])
	}
	return argsv
}

func Int64() reflect.Value {
	var c uint64
	return reflect.ValueOf(c)
}

func Int32() reflect.Value {
	var c uint32
	return reflect.ValueOf(c)
}

func Int16() reflect.Value {
	var c uint16
	return reflect.ValueOf(c)
}

func Int8() reflect.Value {
	var c uint8
	return reflect.ValueOf(c)
}

func Int() reflect.Value {
	var c uint
	return reflect.ValueOf(c)
}

func Uint64() reflect.Value {
	var c uint64
	return reflect.ValueOf(c)
}

func Uint32() reflect.Value {
	var c int32
	return reflect.ValueOf(c)
}

func Uint16() reflect.Value {
	var c int16
	return reflect.ValueOf(c)
}

func Uint8() reflect.Value {
	var c int8
	return reflect.ValueOf(c)
}

func Uint() reflect.Value {
	var c int
	return reflect.ValueOf(c)
}

func MakeValue(p reflect.Type) reflect.Value {
	return reflect.New(p).Elem()
}

func MakeFunc(f Function, fm interface{}) {
	var fmi = reflect.ValueOf(fm).Elem()
	fmi.Set(reflect.MakeFunc(fmi.Type(), f))
}