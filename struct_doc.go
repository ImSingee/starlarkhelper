package starlarkhelper

import (
	"fmt"
	"strings"
)

var _ CanHelp = (*StructDescription)(nil)

func (d *StructDescription) GetDefDoc(mode HelpMode) string {
	var def string
	if d.moduleName != "" {
		def = fmt.Sprintf("Struct [%s.%s]", d.moduleName, d.Name)
	} else {
		def = fmt.Sprintf("Struct [%s]", d.Name)
	}

	switch mode {
	case HelpModeTerminal:
		return terminalBlueString(def)
	default:
		return def
	}
}

func (d *StructDescription) GetSimpleDesc(mode HelpMode) string {
	return strings.ReplaceAll(strings.TrimSpace(d.Short), "\n", "; ")
}

func (d *StructDescription) GetFullDesc(mode HelpMode) string {
	b := strings.Builder{}

	if d.Long != "" {
		b.WriteString(strings.ReplaceAll(strings.TrimSpace(d.Long), "\n", "\n  "))
	}

	if len(d.Fields) != 0 {
		b.WriteString("Fields:\n")
		for _, f := range d.Fields {
			b.WriteString("  - ")
			b.WriteString(ArgGetInDef(f, mode))
			if short := f.DocShort(); short != "" {
				b.WriteString(" ")
				b.WriteString(strings.ReplaceAll(strings.TrimSpace(short), "\n", "; "))
			}
			b.WriteString("\n")
			if long := f.DocLong(); long != "" {
				b.WriteString(" ")
				b.WriteString(strings.ReplaceAll(strings.TrimSpace(long), "\n", "  "))
				b.WriteString("\n")
			}
		}
	}

	return b.String()
}

var _ CanHelp = (*Struct)(nil)

func (s *Struct) GetDefDoc(mode HelpMode) string     { return s.Descriptor.GetDefDoc(mode) }
func (s *Struct) GetSimpleDesc(mode HelpMode) string { return s.Descriptor.GetSimpleDesc(mode) }
func (s *Struct) GetFullDesc(mode HelpMode) string   { return s.Descriptor.GetFullDesc(mode) }
