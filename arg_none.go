package starlarkhelper

type noneArg struct {
	name string
}

func (n noneArg) isArg()                        {}
func (n noneArg) isArgs()                       {}
func (n noneArg) InName() string                { return n.name }
func (n noneArg) InType() string                { return "!" }
func (n noneArg) InDefaultValue() string        { return "" }
func (n noneArg) OutType() string               { return "None" }
func (n noneArg) Optional() Arg                 { return n }
func (n noneArg) Default(value interface{}) Arg { return n }
func (n noneArg) Short(s string) Arg            { return withShort(n, s) }
func (n noneArg) Long(s string) Arg             { return withLong(n, s) }
func (n noneArg) DocShort() string              { return "" }
func (n noneArg) DocLong() string               { return "" }

// NoneArg 标识永远接收/返回 NoneType
func NoneArg(name string) Arg {
	return noneArg{name: name}
}
