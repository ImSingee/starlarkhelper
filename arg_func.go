package starlarkhelper

import (
	"strings"
)

type funcArg struct {
	name string
	in   Args
	out  Arg

	optional   bool
	hasDefault bool

	short string
	long  string
}

func (d *funcArg) isArg()             {}
func (d *funcArg) isArgs()            {}
func (d *funcArg) Optional() Arg      { d.hasDefault = false; d.optional = true; return d }
func (d *funcArg) InName() string     { return d.name }
func (d *funcArg) Short(s string) Arg { d.short = s; return d }
func (d *funcArg) Long(s string) Arg  { d.short = s; return d }

func (d *funcArg) InType() string {
	b := strings.Builder{}

	b.WriteString("f(")
	for i, a := range splitArgs(d.in) {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(a.InType()) // TODO: 暂时只打印类型，name, optional, default 均未打印
	}
	b.WriteString(")")

	if d.out != nil {
		b.WriteString(" -> ")
		b.WriteString(d.out.OutType())
	}

	return b.String()
}
func (d *funcArg) OutType() string { return d.InType() }

func (d *funcArg) Default(value interface{}) Arg { // TODO not-implemented
	d.optional = false
	// not-implemented
	if value != nil {
		d.hasDefault = true
	}

	return d
}

func (d *funcArg) InDefaultValue() string {
	if d.optional {
		return "{}"
	} else if d.hasDefault {
		return "..."
	} else {
		return ""
	}
}

func (d *funcArg) DocShort() string {
	return d.short
}

func (d *funcArg) DocLong() string {
	return d.long
}

func FuncArg(name string, in Args, out Arg) Arg {
	return &funcArg{name: name, in: in, out: out}
}
