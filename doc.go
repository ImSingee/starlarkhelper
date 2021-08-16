package starlarkhelper

import (
	"fmt"
	"strings"
)

// GetArgsInDoc 获取函数入参文档
func (f *Function) GetArgsInDoc() string {
	inArgs := splitArgs(f.ArgsIn)

	if len(inArgs) == 0 {
		// no args
		return ""
	}

	l := make([]string, len(inArgs))
	for i, a := range inArgs {
		b := strings.Builder{}
		b.Grow(64)

		b.WriteString(a.InName())

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

		l[i] = b.String()
	}

	return strings.Join(l, ", ")
}

// GetArgOutDoc 获取函数出参文档
// 返回空或 ` -> TYPE`
func (f *Function) GetArgOutDoc() string {
	if f.ArgOut == nil {
		return ""
	}
	r := f.ArgOut.OutType()
	if r == "" || r == "None" {
		return ""
	} else {
		return " -> " + r
	}
}

func (f *Function) GetDefDoc() string {
	return fmt.Sprintf("%s(%s)%s", f.Name, f.GetArgsInDoc(), f.GetArgOutDoc())
}
