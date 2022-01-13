package allinone

import (
	"reflect"

	"github.com/pkg/errors"
)

func getDummyErr(err error) func() error {
	return func() error { return err }
}

func WrappedFunc(f interface{}, args ...interface{}) func() error {
	if reflect.TypeOf(f).Kind() != reflect.Func {
		err := errors.Errorf("fatal error wrap func want func, got[%s]", reflect.TypeOf(f).Kind())
		//gLogger.Error(err, "wrap func")
		return getDummyErr(err)
	}

	argValues := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValues[i] = reflect.ValueOf(arg)
	}

	return func() error {
		res := reflect.ValueOf(f).Call(argValues)[0]
		return valueAsError(res)
	}
}

func valueAsError(v reflect.Value) error {
	if v.IsNil() {
		return nil
	}
	return v.Interface().(error)
}