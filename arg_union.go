package starlarkhelper

import "strings"

type unionArg struct {
	name         string
	maybe        []Arg
	optional     bool
	defaultValue interface{}

	short string
	long  string
}

func (u *unionArg) isArg()                        {}
func (u *unionArg) isArgs()                       {}
func (u *unionArg) Optional() Arg                 { u.optional = true; return u }
func (u *unionArg) Default(value interface{}) Arg { u.defaultValue = value; return u }

func (u *unionArg) Short(s string) Arg { u.short = s; return u }
func (u *unionArg) Long(s string) Arg  { u.long = s; return u }

func (u *unionArg) InName() string { return u.name }
func (u *unionArg) InType() string {
	maybe := make([]string, len(u.maybe), 1+len(u.maybe))
	for i, m := range u.maybe {
		maybe[i] = m.InType()
	}
	if u.optional {
		maybe = append(maybe, "None")
	}
	return "<" + strings.Join(maybe, " / ") + ">"
}

func (u *unionArg) InDefaultValue() string {
	return defaultValueToString(u.defaultValue)
}

func (u *unionArg) OutType() string {
	return u.InType()
}

func (u *unionArg) DocShort() string {
	return u.short // TODO: 处理 maybe
}

func (u *unionArg) DocLong() string {
	return u.long // TODO: 处理 maybe
}

func UnionArg(name string, a ...Arg) Arg {
	return &unionArg{name: name, maybe: a}
}
