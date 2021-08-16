package starlarkhelper

type listArg struct {
	name  string
	type_ Arg

	optional   bool
	hasDefault bool

	short string
	long  string
}

func (l *listArg) isArg()        {}
func (l *listArg) isArgs()       {}
func (l *listArg) Optional() Arg { l.hasDefault = false; l.optional = true; return l }
func (l *listArg) Default(value interface{}) Arg {
	if value != nil {
		l.optional = false
		l.hasDefault = true
	}
	return l
}

func (l *listArg) Short(s string) Arg { l.short = s; return l }

func (l *listArg) Long(s string) Arg { l.long = s; return l }

func (l *listArg) InName() string { return l.name }

func (l *listArg) InType() string { return "[" + l.type_.InType() + "]" }

func (l *listArg) InDefaultValue() string {
	if l.optional {
		return "[]"
	} else if l.hasDefault {
		return "..."
	} else {
		return ""
	}
}

func (l *listArg) OutType() string { return "[" + l.type_.OutType() + "]" }

func (l *listArg) DocShort() string {
	return l.short // TODO: 处理 maybe
}

func (l *listArg) DocLong() string {
	return l.long // TODO: 处理 maybe
}

func ListArg(name string, innerType Arg) Arg {
	return &listArg{name: name, type_: innerType}
}
