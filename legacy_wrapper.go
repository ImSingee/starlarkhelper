package starlarkhelper

import "go.starlark.net/starlark"

type StandardFunc func(thread *starlark.Thread, fn *Function, args starlark.Tuple, kwargs []starlark.Tuple) (v starlark.Value, err error)

// UnHelp 将 BuiltinFunc 重新映射为标准 StandardFunc，用于兼容第三方库使用
// 需要注意的是，兼容方式要求将原 *starlark.Builtin 类型的 fn 改用 *Function 类型
func UnHelp(f StandardFunc) BuiltinFunc {
	return func(h Helper) (starlark.Value, error) {
		return f(h.Thread, h.Fn, h.PositionalArgs, h.kwargs)
	}
}
