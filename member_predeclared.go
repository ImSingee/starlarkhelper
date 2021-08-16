package starlarkhelper

import (
	"fmt"

	"go.starlark.net/starlark"
)

type PreDeclaredMember struct {
	Name     string
	Short    string    // 简介
	Long     string    // 详细帮助信息
	Examples []Example // 示例

	Value starlark.Value
}

func (m *PreDeclaredMember) getForModule(moduleName string) starlark.Value {
	return m.Value
}

func (m *PreDeclaredMember) Type() string {
	return m.Value.Type()
}

func (m *PreDeclaredMember) GetDefDoc() string {
	return fmt.Sprintf("%s : %s = %s", m.Name, m.Type(), m.Value.String())
}
