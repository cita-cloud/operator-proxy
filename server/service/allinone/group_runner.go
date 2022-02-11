/*
 * Copyright Rivtower Technologies LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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