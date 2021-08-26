package starlarkhelper

import _ "unsafe"

//go:linkname hashString go.starlark.net/starlark.hashString
func hashString(s string) uint32

func Hash(v string, fact uint32) uint32 {
	return hashString(v) ^ fact
}
