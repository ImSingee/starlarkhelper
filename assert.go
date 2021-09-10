package starlarkhelper

import (
	"go.starlark.net/starlark"
)

// Assertable 为数据提供 assert 的能力
type Assertable interface {
	Assert() (pass bool, failReason string)
}

type AssertableBool struct {
	starlark.Bool

	Reason string
}

func (a AssertableBool) Assert() (pass bool, failReason string) {
	return bool(a.Bool), a.Reason
}

func NewFalseWithReason(reason string) AssertableBool {
	return AssertableBool{false, reason}
}
