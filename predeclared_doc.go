package starlarkhelper

import (
	"fmt"
	"strings"
)

var _ CanHelp = (*PreDeclaredValue)(nil)

func (m *PreDeclaredValue) GetDefDoc(mode HelpMode) string {
	def := fmt.Sprintf("%s : %s = %s", m.Name, m.Value.Type(), m.Value.String())

	switch mode {
	case HelpModeTerminal:
		return terminalBlueString(def)
	default:
		return def
	}
}

func (m *PreDeclaredValue) GetSimpleDesc(mode HelpMode) string {
	return strings.TrimSpace(m.Short)
}

func (m *PreDeclaredValue) GetFullDesc(mode HelpMode) string {
	b := strings.Builder{}

	if long := m.Long; long != "" {
		b.WriteString(strings.TrimSpace(long))
		b.WriteString("\n")
	}

	// example
	b.WriteString(GetExamplesDoc(m.Examples, mode))

	return b.String()
}
