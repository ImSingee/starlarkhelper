package starlarkhelper

var _ Arg = (*structArg)(nil)

type structArg struct {
	name string
	desc *StructDescription
}

func (c structArg) isArg()                        {}
func (c structArg) isArgs()                       {}
func (c structArg) Optional() Arg                 { return toOptional(c) }
func (c structArg) Default(value interface{}) Arg { return withDefault(c, value) }
func (c structArg) Short(s string) Arg            { return withShort(c, s) }
func (c structArg) Long(s string) Arg             { return withLong(c, s) }
func (c structArg) InName() string                { return c.name }
func (c structArg) InType() string                { return c.desc.Name }
func (c structArg) InDefaultValue() string        { return "" }
func (c structArg) OutType() string               { return c.InType() }
func (c structArg) DocShort() string              { return "" }
func (c structArg) DocLong() string               { return "" }

func StructArg(name string, desc *StructDescription) Arg {
	return &structArg{name, desc}
}
