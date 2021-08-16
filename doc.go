package starlarkhelper

import (
	"fmt"
	"go.starlark.net/starlark"
	"strings"
)

// CanHelp 接口为自定义类型提供了输出帮助信息的能力
// 内部的 Function Struct Module 已经实现了 CanHelp
type CanHelp interface {
	GetDefDoc() string                // 获取声明信息，应当包括名称、类型，不应包括任何用法解释
	GetSimpleDesc() string            // 获取简单说明信息，应当在一行内完成，不应过长
	GetFullDesc(mode HelpMode) string // 获取详细说明信息，不应重复包括 GetSimpleDesc 中的内容，用于说明具体的使用方式，换行应使用 LF
}

type HelpMode uint8

const (
	HelpModePure     HelpMode = iota // 无格式信息
	HelpModeTerminal                 // 以 terminal 转义 markdown 格式信息
	HelpModeHTML                     // 以 html 转义 markdown 格式信息
)

func title(s string, mode HelpMode) string {
	switch mode {
	case HelpModeTerminal:
		return terminalBlueString(s)
	default:
		return s
	}
}

func tidyAndIndent(s string) string {
	return strings.ReplaceAll(strings.TrimSpace(s), "\n", "  \n")
}

func GetHelpFor(v starlark.Value, mode HelpMode) string {
	if canHelp, ok := v.(CanHelp); ok {
		b := strings.Builder{}

		// 定义
		def := canHelp.GetDefDoc()
		b.WriteString(title(strings.TrimSpace(def), mode))
		b.WriteString("\n")

		// short
		if short := canHelp.GetSimpleDesc(); short != "" {
			b.WriteString(tidyAndIndent(short))
			b.WriteString("\n")
		}

		// long
		if long := canHelp.GetFullDesc(mode); long != "" {
			b.WriteString(tidyAndIndent(long))
			b.WriteString("\n")
		}

		return b.String()
	}

	return getGeneralHelpFor(v)

}

func getGeneralHelpFor(v starlark.Value) string {
	// TODO: 生成更多信息

	return fmt.Sprintf("%s\n%s\n", v.Type(), tidyAndIndent(v.String()))
}

// HelpFunc 返回帮助查看函数
func HelpFunc() *Function {
	return &Function{
		FuncName: "help",
		Short:    "查看帮助信息",
		Long:     "",
		Fn: func(h Helper) (starlark.Value, error) {
			if err := h.CheckExactArgs(1); err != nil {
				return nil, err
			}
			arg, _ := h.GetFirstArg() // err 检查被 CheckExactArgs 涵盖

			h.Print(GetHelpFor(arg, HelpModeTerminal))

			return starlark.None, nil
		},
		ArgsIn: AnyArg("something").Short("要查看帮助的东东"),
		ArgOut: nil, // 输出会直接打印，因此无返回值
	}
}
