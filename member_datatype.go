package starlarkhelper

type DataTypeMember struct {
	Name     string    // 类型名称
	Short    string    // 简介
	Long     string    // 详细帮助信息
	Examples []Example // 示例
}
