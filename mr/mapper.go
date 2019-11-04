package mr

import (
	"github.com/Myriad-Dreamin/functional-go"
	"reflect"
)

type IntMapper struct {
	Mapper
}

func (IntMapper) F(sliceI interface{}) Handler {
	slice := sliceI.([]int)
	handler := make(IntHandler, len(slice), cap(slice))
	copy(handler, slice)
	return handler
}

func (m IntMapper) MapR(fi interface{}) functional.Function {
	if m.MapperF == nil {
		m.MapperF = m
	}
	return m.Mapper.MapR(fi)
}

type MapperTraits struct {
	Mapper
	functional.BaseTraitsInterface
}

func NewMapperTraits(handler interface{}, options ...interface{}) MapperTraits {
	t := MapperTraits{
		BaseTraitsInterface: functional.NewBaseTraits(handler),
	}
	for i := range options {
		switch option := options[i].(type) {
		case int:
			t.CoreCount = option
		}
	}
	return t
}

func (m MapperTraits) F(sliceI interface{}) Handler {
	slice := reflect.ValueOf(sliceI)
	handler := reflect.MakeSlice(m.GetTypeInfo(), slice.Len(), slice.Cap())
	reflect.Copy(handler, slice)
	return handler.Interface().(Handler)
}

func (m MapperTraits) FInplace(slice interface{}) Handler {
	return reflect.ValueOf(slice).Convert(m.GetTypeInfo()).Interface().(Handler)
}

func (m MapperTraits) MapR(fi interface{}) functional.Function {
	if m.MapperF == nil {
		m.MapperF = m
	}
	return m.Mapper.MapR(fi)
}

func (m MapperTraits) Map(fi, fm interface{}) {
	functional.MakeFunc(m.MapR(fi), fm)
}

func (m MapperTraits) MapRInplace(fi interface{}) functional.Function {
	if m.MapperF == nil {
		m.MapperF = m
	}
	return m.Mapper.MapRInplace(fi)
}

func (m MapperTraits) MapInplace(fi, fm interface{}) {
	functional.MakeFunc(m.MapRInplace(fi), fm)
}
