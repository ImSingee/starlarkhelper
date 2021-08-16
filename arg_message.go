package starlarkhelper

type withMessageArg struct {
	Arg

	short string
	long  string
}

func (w *withMessageArg) Short(s string) Arg { w.short = s; return w }

func (w *withMessageArg) Long(l string) Arg { w.long = l; return w }

func (w *withMessageArg) DocShort() string { return w.short }

func (w *withMessageArg) DocLong() string { return w.long }

func withMessage(arg Arg) Arg {
	return &withMessageArg{Arg: arg}
}

func withShort(arg Arg, s string) Arg {
	return &withMessageArg{Arg: arg, short: s}
}

func withLong(arg Arg, l string) Arg {
	return &withMessageArg{Arg: arg, long: l}
}
