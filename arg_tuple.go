package starlarkhelper

import (
	"fmt"
	"strings"
)

type namedTupleArg struct {
	name         string
	contains     []Arg
	optional     bool
	defaultValue interface{}

	short string
	long  string
}

func (u *namedTupleArg) isArg()                        {}
func (u *namedTupleArg) isArgs()                       {}
func (u *namedTupleArg) Optional() Arg                 { u.optional = true; return u }
func (u *namedTupleArg) Default(value interface{}) Arg { u.defaultValue = value; return u }

func (u *namedTupleArg) Short(s string) Arg { u.short = s; return u }
func (u *namedTupleArg) Long(s string) Arg  { u.long = s; return u }

func (u *namedTupleArg) InName() string { return u.name }
func (u *namedTupleArg) InType() string {
	maybe := make([]string, len(u.contains))
	for i, m := range u.contains {
		maybe[i] = fmt.Sprintf("%s: %s", m.InName(), m.InType())
	}

	suffix := ""
	if u.optional {
		suffix = "?"
	}
	return "(" + strings.Join(maybe, ", ") + ")" + suffix
}

func (u *namedTupleArg) InDefaultValue() string {
	return defaultValueToString(u.defaultValue)
}

func (u *namedTupleArg) OutType() string {
	return u.InType()
}

func (u *namedTupleArg) DocShort() string {
	return u.short // TODO: 处理 contains
}

func (u *namedTupleArg) DocLong() string {
	return u.long // TODO: 处理 contains
}

func TupleArg(name string, a ...Arg) Arg {
	return &namedTupleArg{name: name, contains: a}
}
