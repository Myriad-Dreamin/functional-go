package functional

import (
	"reflect"
)

func AddR(in []reflect.Value) []reflect.Value {
	var leftType = in[0].Type()
	if leftType == in[1].Type() {
		c := MakeValue(leftType)
		switch {
		case IsSigned(leftType):
			a, b := in[0].Int(), in[1].Int()
			c.Set(reflect.ValueOf(a+b).Convert(c.Type()))
		case IsUnsigned(leftType):
			a, b := in[0].Uint(), in[1].Uint()
			c.Set(reflect.ValueOf(a+b).Convert(c.Type()))
		case IsFloat(leftType):
			a, b := in[0].Float(), in[1].Float()
			c.Set(reflect.ValueOf(a+b).Convert(c.Type()))
		}
		return []reflect.Value{c}
	}
	panic("")
}

func Add(f interface{}) {
	var fi = reflect.ValueOf(f).Elem()
	fi.Set(reflect.MakeFunc(fi.Type(), AddR))
}

func AddRC(in []reflect.Value) []reflect.Value {
	var leftType = in[0].Type()
	if leftType == in[1].Type() {
		c := in[2].Elem()
		switch {
		case IsSigned(leftType):
			a, b := in[0].Int(), in[1].Int()
			c.Set(reflect.ValueOf(a+b).Convert(c.Type()))
		case IsUnsigned(leftType):
			a, b := in[0].Uint(), in[1].Uint()
			c.Set(reflect.ValueOf(a+b).Convert(c.Type()))
		case IsFloat(leftType):
			a, b := in[0].Float(), in[1].Float()
			c.Set(reflect.ValueOf(a+b).Convert(c.Type()))
		}
		return nil
	}
	panic("")
}

func AddC(f interface{}) {
	var fi = reflect.ValueOf(f).Elem()
	fi.Set(reflect.MakeFunc(fi.Type(), func(in []reflect.Value) []reflect.Value {
		c := reflect.New(fi.Type().Out(0))
		AddRC(append(in, c))
		return []reflect.Value{c.Elem()}
	}))
}
