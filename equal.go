package starlarkhelper

import (
	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

func Equal(one, another starlark.Value) bool {
	ok, err := starlark.Compare(syntax.EQL, one, another)
	if err != nil {
		return false
	}
	return ok
}

func EqualInt(one starlark.Value, another int64) bool {
	return Equal(one, starlark.MakeInt64(another))
}
