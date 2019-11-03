package functional

import "reflect"


func IsSigned(i reflect.Type) bool {
	switch i.Kind() {
	case reflect.Int64,reflect.Int32,reflect.Int16,reflect.Int8,reflect.Int:
		return true
	}
	return false
}

func IsUnsigned(i reflect.Type) bool {
	switch i.Kind() {
	case reflect.Uint64,reflect.Uint32,reflect.Uint16,reflect.Uint8,reflect.Uint:
		return true
	}
	return false
}

func IsFloat(i reflect.Type) bool {
	switch i.Kind() {
	case reflect.Float64, reflect.Float32:
		return true

	}
	return false
}

func IsInteger(i reflect.Type) bool {
	switch i.Kind() {
	case reflect.Uint64,reflect.Uint32,reflect.Uint16,reflect.Uint8,reflect.Uint,
		reflect.Int64,reflect.Int32,reflect.Int16,reflect.Int8,reflect.Int:
		return true
	}
	return false
}

func IsNumeric(i reflect.Type) bool {
	switch i.Kind() {
	case reflect.Uint64,reflect.Uint32,reflect.Uint16,reflect.Uint8,reflect.Uint,
		reflect.Int64,reflect.Int32,reflect.Int16,reflect.Int8,reflect.Int,
		reflect.Float64, reflect.Float32:
		return true
	}
	return false
}

