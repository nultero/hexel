package hexel

import (
	"fmt"
	"reflect"
	"runtime"
)

func blame(function interface{}, err error) {
	funcName := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	msg := fmt.Sprintf(
		"func %v has thrown unrecoverable err: %v",
		funcName, err,
	)
	panic(msg)
}
