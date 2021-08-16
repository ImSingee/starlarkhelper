package starlarkhelper

import (
	"fmt"
	"strings"
)

// dictArg
type dictArg struct {
	name      string
	keyType   Arg
	valueType Arg

	optional   bool
	hasDefault bool

	short string
	long  string
}

func (d *dictArg) isArg()             {}
func (d *dictArg) isArgs()            {}
func (d *dictArg) Optional() Arg      { d.hasDefault = false; d.optional = true; return d }
func (d *dictArg) InName() string     { return d.name }
func (d *dictArg) Short(s string) Arg { d.short = s; return d }
func (d *dictArg) Long(s string) Arg  { d.short = s; return d }

func (d *dictArg) InType() string {
	return fmt.Sprintf("{[%s: %s]: %s: %s}",
		d.keyType.InName(), d.keyType.InType(),
		d.valueType.InName(), d.valueType.InType(),
	)
}
func (d *dictArg) OutType() string {
	return fmt.Sprintf("dict<key = %s, value = %s>", d.keyType.OutType(), d.valueType.OutType())
}

func (d *dictArg) Default(value interface{}) Arg {
	d.optional = false
	// not-implemented
	if value != nil {
		d.hasDefault = true
	}

	return d
}

func (d *dictArg) InDefaultValue() string {
	if d.optional {
		return "{}"
	} else if d.hasDefault {
		return "..."
	} else {
		return ""
	}
}

func mergeShort(k, v string) string {
	if k == "" && v == "" {
		return ""
	}
	if k != "" && v != "" {
		return fmt.Sprintf("Key: %s, Value: %s", k, v)
	}

	if k != "" {
		return fmt.Sprintf("Key: %s", k)
	} else { //  v != ""
		return fmt.Sprintf("Value: %s", v)
	}
}

func mergeLong(k, v string) string {
	if k == "" && v == "" {
		return ""
	}

	b := strings.Builder{}

	if k != "" {
		b.WriteString("Key:\n  ")
		b.WriteString(strings.ReplaceAll(strings.TrimSpace(k), "\n", "\n  "))
	}

	if v != "" {
		b.WriteString("Value:\n  ")
		b.WriteString(strings.ReplaceAll(strings.TrimSpace(v), "\n", "\n  "))
	}

	return b.String()
}

func (d *dictArg) DocShort() string {
	if d.short != "" {
		return d.short
	} else {
		return mergeShort(d.keyType.DocShort(), d.valueType.DocShort())
	}
}

func (d *dictArg) DocLong() string {
	if d.long != "" {
		return d.long
	} else {
		return mergeLong(d.keyType.DocLong(), d.valueType.DocLong())
	}
}

func DictArg(name string, keyType, valueType Arg) Arg {
	return &dictArg{name: name, keyType: keyType, valueType: valueType}
}
