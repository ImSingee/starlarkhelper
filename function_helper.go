package starlarkhelper

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/pkg/errors"
)

type BuiltinFunc func(h Helper) (starlark.Value, error)
type UnpackArgsFunc func(pairs ...interface{}) error

type Helper struct {
	Name   string
	Thread *starlark.Thread
	Fn     *Function
	Args   starlark.Tuple
	Kwargs []starlark.Tuple

	Err error
}

func (h *Helper) Print(msg string) {
	h.Thread.Print(h.Thread, msg)
}

func (h *Helper) withErrorHandler(f func() error) *Helper {
	if h.Err == nil {
		h.Err = f()
	}
	return h
}

func (h *Helper) UnpackArgs(pairs ...interface{}) *Helper {
	return h.withErrorHandler(func() error {
		return starlark.UnpackArgs(h.Name, h.Args, h.Kwargs, pairs...)
	})
}

// UnpackArgsIgnoreKeyword 类似于 UnpackArgs，但是忽略命名参数
func (h *Helper) UnpackArgsIgnoreKeyword(pairs ...interface{}) *Helper {
	return h.withErrorHandler(func() error {
		return starlark.UnpackArgs(h.Name, h.Args, nil, pairs...)
	})
}

func (h *Helper) UnpackPositionalArgs(min int, vars ...interface{}) *Helper {
	return h.withErrorHandler(func() error {
		return starlark.UnpackPositionalArgs(h.Name, h.Args, h.Kwargs, min, vars...)
	})
}

// UnpackBasicArgs 类似于 UnpackArgs, 但仅支持 int, int64, uint64, string, bool
func (h *Helper) UnpackBasicArgs(pairs ...interface{}) *Helper {
	return h.withErrorHandler(func() error {
		ppp := make([][3]interface{}, len(pairs)/2)
		for i := 0; i < len(ppp); i++ {
			ppp[i][0] = pairs[i*2]   // 名称
			ppp[i][1] = pairs[i*2+1] // 指向基本类型的指针
			switch v := ppp[i][1].(type) {
			case *int:
				ppp[i][2] = starlark.MakeInt(*v)
			case *int64:
				ppp[i][2] = starlark.MakeInt64(*v)
			case *uint64:
				ppp[i][2] = starlark.MakeUint64(*v)
			case *string:
				ppp[i][2] = starlark.String(*v)
			case *bool:
				ppp[i][2] = starlark.Bool(*v)
			default:
				return errors.Errorf("unsupported type %T", v)
			}
			pairs[i*2+1] = &ppp[i][2] // 指向 starlark 类型的指针
		}

		err := starlark.UnpackArgs(h.Name, h.Args, h.Kwargs, pairs...)
		if err != nil {
			return err
		}
		for i := 0; i < len(ppp); i++ {
			switch v := ppp[i][1].(type) {
			case *int:
				vv, ok := ppp[i][2].(starlark.Int).Int64()
				if !ok {
					return errors.Errorf("arg %s: overflow int64", ppp[i][0])
				}
				*v = int(vv)
			case *int64:
				vv, ok := ppp[i][2].(starlark.Int).Int64()
				if !ok {
					return errors.Errorf("arg %s: overflow int64", ppp[i][0])
				}
				*v = vv
			case *uint64:
				vv, ok := ppp[i][2].(starlark.Int).Uint64()
				if !ok {
					return errors.Errorf("arg %s: overflow uint64", ppp[i][0])
				}
				*v = vv
			case *string:
				*v = string(ppp[i][2].(starlark.String))
			case *bool:
				*v = bool(ppp[i][2].(starlark.Bool))
			default: // impossible
			}
		}

		return nil
	})
}

// GetKeywordArgs 从 keyword args 中提取出指定 key 对应的值，如果值不存在返回 nil
func (h *Helper) GetKeywordArgs(key string) starlark.Value {
	return GetFromArgs(h.Kwargs, key)
}

// GetKeywordArgsBool 从 keyword args 中提取出指定 key 对应的值
func (h *Helper) GetKeywordArgsBool(key string, defaultValue bool) (bool, error) {
	vv := GetFromArgs(h.Kwargs, key)
	if vv == nil {
		return defaultValue, nil
	}

	if v, ok := vv.(starlark.Bool); ok {
		return bool(v), nil
	}

	return false, fmt.Errorf("arg %s: invalid type (not bool)", key)
}

func (h *Helper) ConvertString(convert starlark.String, to *string) *Helper {
	*to = string(convert)
	return h
}

// ArgsCount 返回传入的 args 数量
func (h *Helper) ArgsCount() int {
	return h.Args.Len() + len(h.Kwargs)
}

// CheckExactArgs 检查是否传递了指定数量的参数，如果否则返回 error
func (h *Helper) CheckExactArgs(count int) error {
	c := h.ArgsCount()

	if c != count {
		return fmt.Errorf("expect %d args, got %d", count, c)
	}

	return nil
}

// CheckMinArgs 检查是否传递了至少指定数量的参数，如果否则返回 error
func (h *Helper) CheckMinArgs(count int) error {
	c := h.ArgsCount()

	if c < count {
		return fmt.Errorf("expect at least %d args, got %d", count, c)
	}

	return nil
}

// GetFirstArg 获得第一个参数
func (h *Helper) GetFirstArg() (starlark.Value, error) {
	if err := h.CheckMinArgs(1); err != nil {
		return nil, err
	}

	if h.Args.Len() > 0 {
		return h.Args.Index(0), nil
	} else {
		return h.Kwargs[0].Index(1), nil
	}
}

func (h *Helper) GetAllPositionalArgs() []starlark.Value {
	return []starlark.Value(h.Args)
}
