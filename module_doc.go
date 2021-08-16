package starlarkhelper

import (
	"fmt"
	"strings"
)

var _ CanHelp = (*StarlarkModule)(nil)

func (m *StarlarkModule) GetDefDoc(mode HelpMode) string {
	def := fmt.Sprintf("Module [%s]", m.Name)

	switch mode {
	case HelpModeTerminal:
		return terminalCyanString(def)
	default:
		return def
	}
}

func (m *StarlarkModule) GetSimpleDesc(mode HelpMode) string {
	return m.internal.Short
}

func (m *StarlarkModule) GetFullDesc(mode HelpMode) string {
	b := strings.Builder{}

	for _, member := range m.internal.PreDeclares {
		b.WriteString(terminalBlueString("[预定义变量] "))
		b.WriteString(strings.ReplaceAll(strings.TrimSpace(GetHelperHelpFor(member, mode)), "\n", "\n  "))
		b.WriteString("\n")
	}

	for _, member := range m.internal.Functions {
		b.WriteString(terminalBlueString("[函数] "))
		b.WriteString(strings.ReplaceAll(strings.TrimSpace(GetHelperHelpFor(member, mode)), "\n", "\n  "))
		b.WriteString("\n")
	}

	return b.String()

}
