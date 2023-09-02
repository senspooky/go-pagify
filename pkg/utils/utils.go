package utils

import "reflect"

func GetPointerToInterface(v interface{}) interface{} {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Pointer {
		if tt := t.Elem(); tt.Kind() == reflect.Pointer {
			panic("does not support pointer to interface")
		} else {
			if reflect.ValueOf(v).IsZero() {
				vv := reflect.New(tt)
				return vv.Interface()
			}
		}

	}
	return reflect.New(t).Interface()
}

func GetUnderlyingType(v interface{}) reflect.Type {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}
