package starlarkhelper

import (
	"context"
	"fmt"
	"go.starlark.net/starlark"
)

// Function 函数类型
type Function struct {
	Ctx context.Context

	FuncName string // 函数名称
	Short    string // 简介
	Long     string // 详细帮助信息

	Examples []Example // 示例

	Middleware Middleware
	Fn         BuiltinFunc
	ArgsIn     Args
	ArgOut     Arg

	recv       starlark.Value // 接收者，来源于 Struct 时会用到
	moduleName string         // 模块名称，来源于 Module 时会用到
}

var _ starlark.Callable = (*Function)(nil)

// NewBuiltinFunction 快速创建一个无文档的 Function 示例
func NewBuiltinFunction(name string, f BuiltinFunc) *Function {
	return &Function{
		Ctx:      context.TODO(),
		FuncName: name,
		Fn:       f,
	}
}

// Name 为增加了 module 名称的 name
func (f *Function) Name() string {
	if f.moduleName != "" {
		return f.moduleName + "." + f.FuncName
	} else {
		return f.FuncName
	}
}

func (f *Function) String() string           { return fmt.Sprintf("<built-in function %s>", f.Name()) }
func (f *Function) Type() string             { return "builtin_function_or_method" }
func (f *Function) Receiver() starlark.Value { return f.recv }
func (f *Function) Truth() starlark.Bool     { return true }
func (f *Function) Hash() (uint32, error)    { return Hash(f.Name(), 3212), nil }

func (f *Function) Freeze() {
	if f.recv != nil {
		f.recv.Freeze()
	}
}

func (f *Function) CallInternal(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	ctx := f.Ctx
	if ctx == nil {
		ctx = context.TODO()
	}

	return f.call(ctx, &Helper{
		Name:           f.Name(),
		Thread:         thread,
		Fn:             f,
		PositionalArgs: []starlark.Value(args),
		KeywordArgs:    convertKwargsToStringDict(kwargs),
		Args:           args,
		Kwargs:         kwargs,
	})
}

func (f *Function) call(ctx context.Context, h *Helper) (starlark.Value, error) {
	if f.Middleware == nil {
		return f.Fn(ctx, h)
	} else {
		return f.Middleware(ctx, h, f.Fn)
	}
}

func convertKwargsToStringDict(kwargs []starlark.Tuple) starlark.StringDict {
	d := make(starlark.StringDict, len(kwargs))
	for _, pair := range kwargs {
		d[string(pair[0].(starlark.String))] = pair[1]
	}
	return d
}

func (f Function) Copy() *Function {
	return &f
}
