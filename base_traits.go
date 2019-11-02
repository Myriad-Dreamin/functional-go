package functional

import "reflect"

type BaseTraitsInterface interface {
	GetTypeInfo() reflect.Type
	GetSliceInfo() reflect.Type
	ObjectFactory() interface{}
	SliceFactory() interface{}
}

type BaseTraits struct {
	TypeInfo reflect.Type
	SliceInfo reflect.Type
}

func (traits BaseTraits) ObjectFactory() interface{} {
	return reflect.New(traits.TypeInfo).Interface()
}

func (traits BaseTraits) SliceFactory() interface{} {
	return reflect.MakeSlice(traits.SliceInfo, 0, 0).Interface()
}

func (traits BaseTraits) GetTypeInfo() reflect.Type {
	return traits.TypeInfo
}

func (traits BaseTraits) GetSliceInfo() reflect.Type {
	return traits.SliceInfo
}


func NewBaseTraits(t interface{}) BaseTraits {
	traits := BaseTraits{}
	traits.TypeInfo = reflect.TypeOf(t)
	traits.SliceInfo = reflect.SliceOf(traits.TypeInfo)
	return traits
}