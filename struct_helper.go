package starlarkhelper

import "go.starlark.net/starlark"

func (s *Struct) FieldEqual(name string, value starlark.Value) bool {
	return Equal(s.Values[name], value)
}
