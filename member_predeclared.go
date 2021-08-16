package starlarkhelper

import (
	"fmt"

	"go.starlark.net/starlark"
)

// PreDeclaredValue 为预声明的变量/常量值
type PreDeclaredValue struct {
	Name     string
	Short    string    // 简介
	Long     string    // 详细帮助信息
	Examples []Example // 示例

	Value starlark.Value
}

func (m *PreDeclaredValue) getForModule(moduleName string) starlark.Value {
	return m.Value
}

func (m *PreDeclaredValue) Type() string {
	return m.Value.Type()
}

func (m *PreDeclaredValue) GetDefDoc() string {
	return fmt.Sprintf("%s : %s = %s", m.Name, m.Type(), m.Value.String())
}
