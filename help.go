package starlarkhelper

// CanHelp 接口为自定义类型提供了输出帮助信息的能力
// 内部的 Function Struct Module 已经实现了 CanHelp
type CanHelp interface {
	Help() string
}

func HelpFunc() {

}
