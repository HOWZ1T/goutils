package reflect

import (
	"errors"
	"reflect"
	"runtime"
)

func IsFunction(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

// causes a panic when a non-function argument is passed as v
func GetFunctionName(v interface{}) string {
	if IsFunction(v) {
		return runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
	}

	panic(errors.New("GetFunctionName: cannot get the name of passed parameter v as it is NOT a function"))
}
