package starlarkhelper

type Args interface {
	isArgs()
}

var _ Args = (Arg)(nil)

var _ Args = (multipleArgs)(nil)

type multipleArgs []Arg

func (multipleArgs) isArgs() {}

// M 标识入参为多个
func M(a ...Arg) Args {
	return multipleArgs(a)
}

func splitArgs(args Args) []Arg {
	switch a := args.(type) {
	case nil:
		return nil
	case Arg:
		return []Arg{a}
	default:
		return []Arg(args.(multipleArgs))
	}
}
