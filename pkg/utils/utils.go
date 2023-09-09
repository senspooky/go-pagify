package utils

import (
	"fmt"
	"reflect"
)

func CopyInterfaceValues(src, dst interface{}) (err error) {
	defer func() { // method contains reflections, so it may panic
		if p := recover(); p != nil {
			err = fmt.Errorf("unable to copy interface values: %v", p)
		}
	}()
	reflect.Indirect(reflect.ValueOf(dst)).Set(reflect.Indirect(reflect.ValueOf(src)))
	return
}

func GetNewPointerToInterface(v interface{}) (ptr interface{}, err error) {
	defer func() { // method contains reflections, so it may panic
		if p := recover(); p != nil {
			err = fmt.Errorf("unable to get pointer: %v", p)
		}
	}()

	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Pointer {
		if tt := t.Elem(); tt.Kind() == reflect.Pointer {
			panic("does not support pointer to interface")
		} else {
			if reflect.ValueOf(v).IsZero() {
				vv := reflect.New(tt)
				return vv.Interface(), nil
			}
		}

	}
	return reflect.New(t).Interface(), nil
}

func CompareTyes(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}
