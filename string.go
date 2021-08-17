package starlarkhelper

import (
	"fmt"
	"go.starlark.net/starlark"
)

// String convert anything to string
func String(any fmt.Stringer) starlark.String {
	return starlark.String(any.String())
}
