package starlarkhelper

import _ "unsafe"

//go:linkname hashString go.starlark.net/starlark.hashString
func hashString(s string) uint32
