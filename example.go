package starlarkhelper

import "strings"

type Example struct {
	Comment string
	Code    string
}

func (e *Example) GetDoc(mode HelpMode) string {
	b := strings.Builder{}

	switch mode {
	case HelpModeTerminal:
		b.WriteString(terminalUnderlineString("Example: "))
	default:
		b.WriteString("Example: ")
	}
	b.WriteString(strings.ReplaceAll(strings.TrimSpace(e.Comment), "\n", "; "))
	b.WriteString("\n")

	for _, line := range strings.Split(strings.TrimSpace(e.Code), "\n") {
		b.WriteString("  ")
		b.WriteString(terminalItalicString(line))
		b.WriteString("\n")
	}

	return b.String()
}

func GetExamplesDoc(examples []Example, mode HelpMode) string {
	if examples := examples; len(examples) != 0 {
		result := make([]string, len(examples))

		for i, example := range examples {
			result[i] = example.GetDoc(mode)
		}

		return strings.Join(result, "")
	} else {
		return ""
	}
}
