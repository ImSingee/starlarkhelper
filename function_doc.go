package starlarkhelper

import (
	"fmt"
	"strings"
)

var _ CanHelp = (*Function)(nil)

func (f *Function) GetDefDoc(mode HelpMode) string {
	def := fmt.Sprintf("%s(%s)%s", f.Name(), f.GetArgsInDef(), f.GetArgOutDef())

	switch mode {
	case HelpModeTerminal:
		return terminalBlueString(def)
	default:
		return def
	}
}

func (f *Function) GetSimpleDesc(mode HelpMode) string {
	return f.Short
}

func (f *Function) GetFullDesc(mode HelpMode) string {
	b := strings.Builder{}

	// 入参信息
	b.WriteString(" In:")
	if f.ArgsIn == ELLIPSIS {
		b.WriteString(" ...\n")
	} else if inArgs := f.ArgsInSplit(); len(inArgs) == 0 {
		b.WriteString(" (no args)\n")
	} else {
		b.WriteString("\n")
		for _, a := range inArgs {
			if _, ok := a.(symbolArg); ok {
				continue
			}
			b.WriteString("  - ")
			b.WriteString(ArgGetInDef(a, mode))
			if short := a.DocShort(); short != "" {
				b.WriteString(" ")
				b.WriteString(strings.ReplaceAll(strings.TrimSpace(short), "\n", "; "))
			}
			b.WriteString("\n")
			if long := a.DocLong(); long != "" {
				b.WriteString("      ")
				b.WriteString(strings.ReplaceAll(strings.TrimSpace(long), "\n", "\n      "))
				b.WriteString("\n")
			}
		}
	}

	// 出参信息
	b.WriteString("Out: ")
	if outArg := f.ArgOut; outArg == ELLIPSIS {
		b.WriteString("...")
	} else if outArg == nil {
		b.WriteString("None")
	} else {
		b.WriteString(ArgGetOutDef(f.ArgOut))

		if short := f.ArgOut.DocShort(); short != "" {
			b.WriteString("  ")
			b.WriteString(strings.ReplaceAll(strings.TrimSpace(short), "\n", "; "))
			b.WriteString("\n")
		}
		if long := f.ArgOut.DocLong(); long != "" {
			b.WriteString("  ")
			b.WriteString(strings.ReplaceAll(strings.TrimSpace(long), "\n", "\n  "))
			b.WriteString("\n")
		}

	}
	b.WriteString("\n")

	// 更多描述
	if long := f.Long; long != "" {
		b.WriteString(strings.TrimSpace(f.Long))
		b.WriteString("\n")
	}

	// example
	b.WriteString(GetExamplesDoc(f.Examples, mode))

	return b.String()
}

func (f *Function) ArgsInSplit() []Arg {
	return splitArgs(f.ArgsIn)
}

// GetArgsInDef 获取函数入参文档
func (f *Function) GetArgsInDef() string {
	inArgs := f.ArgsInSplit()

	if len(inArgs) == 0 {
		// no args
		return ""
	}

	l := make([]string, len(inArgs))
	for i, a := range inArgs {
		l[i] = ArgGetInDef(a, HelpModePure)
	}

	return strings.Join(l, ", ")
}

// GetArgOutDef 获取函数出参文档
// 返回空或 ` -> TYPE`
func (f *Function) GetArgOutDef() string {
	o := ArgGetOutDef(f.ArgOut)
	if o == "" {
		return o
	} else {
		return " -> " + o
	}
}
