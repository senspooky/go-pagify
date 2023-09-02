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
