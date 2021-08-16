package starlarkhelper

import "go.starlark.net/starlark"

type Member interface {
	getForModule(moduleName string) StarlarkMember
}

type StarlarkMember interface {
	starlark.Value

	Internal() Member
}

type StarlarkBuiltinMember struct {
	*starlark.Builtin

	internal Member
}

func (m StarlarkBuiltinMember) Internal() Member { return m.internal }
