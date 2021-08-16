package starlarkhelper

import "go.starlark.net/starlark"

var _ Arg = (*CustomTypeArg)(nil)

type CustomTypeArg struct {
	Name string
	Type starlark.Value
}

func (c CustomTypeArg) isArg()                        {}
func (c CustomTypeArg) isArgs()                       {}
func (c CustomTypeArg) Optional() Arg                 { return toOptional(c) }
func (c CustomTypeArg) Default(value interface{}) Arg { return withDefault(c, value) }
func (c CustomTypeArg) Short(s string) Arg            { return withShort(c, s) }
func (c CustomTypeArg) Long(s string) Arg             { return withLong(c, s) }
func (c CustomTypeArg) InName() string                { return c.Name }
func (c CustomTypeArg) InType() string                { return c.Type.Type() }
func (c CustomTypeArg) InDefaultValue() string        { return "" }
func (c CustomTypeArg) OutType() string               { return c.InType() }
func (c CustomTypeArg) DocShort() string              { return "" }
func (c CustomTypeArg) DocLong() string               { return "" }
