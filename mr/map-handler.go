package mr

import (
	"reflect"
)

type InterfaceHandler []interface{}

type InterfaceMapFunc = func(interface{}) interface{}

func (handler InterfaceHandler) Call(f interface{}, index int) {
	handler[index] = f.(InterfaceMapFunc)(handler[index])
}

func (handler InterfaceHandler) Value() reflect.Value {
	return reflect.ValueOf(([]interface{})(handler))
}

type ReflectHandler []reflect.Value

type ReflectMapFunc = func(reflect.Value) reflect.Value

func (handler ReflectHandler) Call(f interface{}, index int) {
	handler[index] = f.(ReflectMapFunc)(handler[index])
}

func (handler ReflectHandler) Value() reflect.Value {
	return reflect.ValueOf(([]reflect.Value)(handler))
}

type IntHandler []int

type IntMapFunc = func(int) int

func (handler IntHandler) Call(f interface{}, index int) {
	handler[index] = f.(IntMapFunc)(handler[index])
}

func (handler IntHandler) Value() reflect.Value {
	return reflect.ValueOf(([]int)(handler))
}

type Int8Handler []int8

type Int8MapFunc = func(int8) int8

func (handler Int8Handler) Call(f interface{}, index int) {
	handler[index] = f.(Int8MapFunc)(handler[index])
}

func (handler Int8Handler) Value() reflect.Value {
	return reflect.ValueOf(([]int8)(handler))
}

type Int16Handler []int16

type Int16MapFunc = func(int16) int16

func (handler Int16Handler) Call(f interface{}, index int) {
	handler[index] = f.(Int16MapFunc)(handler[index])
}

func (handler Int16Handler) Value() reflect.Value {
	return reflect.ValueOf(([]int16)(handler))
}

type Int32Handler []int32

type Int32MapFunc = func(int32) int32

func (handler Int32Handler) Call(f interface{}, index int) {
	handler[index] = f.(Int32MapFunc)(handler[index])
}

func (handler Int32Handler) Value() reflect.Value {
	return reflect.ValueOf(([]int32)(handler))
}


type Int64Handler []int64

type Int64MapFunc = func(int64) int64

func (handler Int64Handler) Call(f interface{}, index int) {
	handler[index] = f.(Int64MapFunc)(handler[index])
}

func (handler Int64Handler) Value() reflect.Value {
	return reflect.ValueOf(([]int64)(handler))
}

type UintHandler []uint

type UintMapFunc = func(uint) uint

func (handler UintHandler) Call(f interface{}, index int) {
	handler[index] = f.(UintMapFunc)(handler[index])
}

func (handler UintHandler) Value() reflect.Value {
	return reflect.ValueOf(([]uint)(handler))
}

type Uint8Handler []uint8

type Uint8MapFunc = func(uint8) uint8

func (handler Uint8Handler) Call(f interface{}, index int) {
	handler[index] = f.(Uint8MapFunc)(handler[index])
}

func (handler Uint8Handler) Value() reflect.Value {
	return reflect.ValueOf(([]uint8)(handler))
}

type Uint16Handler []uint16

type Uint16MapFunc = func(uint16) uint16

func (handler Uint16Handler) Call(f interface{}, index int) {
	handler[index] = f.(Uint16MapFunc)(handler[index])
}

func (handler Uint16Handler) Value() reflect.Value {
	return reflect.ValueOf(([]uint16)(handler))
}

type Uint32Handler []uint32

type Uint32MapFunc = func(uint32) uint32

func (handler Uint32Handler) Call(f interface{}, index int) {
	handler[index] = f.(Uint32MapFunc)(handler[index])
}

func (handler Uint32Handler) Value() reflect.Value {
	return reflect.ValueOf(([]uint32)(handler))
}

type Uint64Handler []uint64

type Uint64MapFunc = func(uint64) uint64

func (handler Uint64Handler) Call(f interface{}, index int) {
	handler[index] = f.(Uint64MapFunc)(handler[index])
}

func (handler Uint64Handler) Value() reflect.Value {
	return reflect.ValueOf(([]uint64)(handler))
}


