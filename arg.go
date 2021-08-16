package starlarkhelper

import "strings"

type Arg interface {
	isArg()
	isArgs()

	Optional() Arg // 将 arg 标记为 optional
	Default(value interface{}) Arg
	Short(s string) Arg
	Long(s string) Arg

	InName() string         // 传入时的参数名
	InType() string         // 传入时展示的参数名
	InDefaultValue() string // 传入时的默认值（空代表无默认值）
	OutType() string        // 返回时展示的参数名（空代表无返回值）

	DocShort() string
	DocLong() string
}

func ArgGetInDef(a Arg, mode HelpMode) string {
	if a == nil {
		return ""
	}

	b := strings.Builder{}
	b.Grow(64)

	switch mode {
	case HelpModeTerminal:
		b.WriteString(terminalBoldString(a.InName()))
	default:
		b.WriteString(a.InName())
	}

	t := a.InType()
	d := a.InDefaultValue()

	if t != "" {
		b.WriteString(": ")
		b.WriteString(t)
	}
	if d != "" {
		b.WriteString(" = ")
		b.WriteString(d)
	}

	return b.String()
}

func ArgGetOutDef(a Arg) string {
	if a == nil {
		return ""
	}
	r := a.OutType()
	if r == "" || r == "None" {
		return ""
	} else {
		return r
	}
}
