package starlarkhelper

type optionalArg struct {
	Arg
}

func (n optionalArg) InType() string         { return n.Arg.InType() + "?" }
func (n optionalArg) InDefaultValue() string { return "" }
func (n optionalArg) OutType() string        { return "<" + n.Arg.OutType() + " / None>" }
func (n optionalArg) Optional() Arg          { return n }

// toOptional 标识可选类型
func toOptional(arg Arg) Arg {
	return optionalArg{Arg: arg}
}
