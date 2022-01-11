package starlarkhelper

import _ "unsafe"

// hashString computes the hash of s.
// the function copied from go.starlark.net/starlark.hashString
func hashString(s string) uint32 {
	if len(s) >= 12 {
		// Call the Go runtime's optimized hash implementation,
		// which uses the AESENC instruction on amd64 machines.
		return uint32(goStringHash(s, 0))
	}
	return softHashString(s)
}

//go:linkname goStringHash runtime.stringHash
func goStringHash(s string, seed uintptr) uintptr

// softHashString computes the 32-bit FNV-1a hash of s in software.
// the function copied from go.starlark.net/starlark.softHashString
func softHashString(s string) uint32 {
	var h uint32 = 2166136261
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

func Hash(v string, fact uint32) uint32 {
	return hashString(v) ^ fact
}
