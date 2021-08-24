package starlarkhelper

import (
	"github.com/pkg/errors"
	"go.starlark.net/starlark"
)

func AsEvalError(err error) *starlark.EvalError {
	var e *starlark.EvalError
	if errors.As(err, &e) {
		return e
	} else {
		return nil
	}
}

func CallStack(err error) starlark.CallStack {
	e := AsEvalError(err)
	if e == nil {
		return nil
	}
	return e.CallStack
}
