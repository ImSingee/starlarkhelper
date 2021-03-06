package starlarkhelper

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"go.starlark.net/starlark"
)

type BuiltinFunc func(ctx context.Context, h *Helper) (starlark.Value, error)
type UnpackArgsFunc func(pairs ...interface{}) error

type Helper struct {
	Name           string
	Thread         *starlark.Thread
	Fn             *Function
	PositionalArgs []starlark.Value
	Kwargs         []starlark.Tuple // original Kwargs
	KeywordArgs    starlark.StringDict

	Err error

	Args starlark.Tuple // Deprecated: please use PositionalArgs
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
		return starlark.UnpackArgs(h.Name, h.PositionalArgs, h.Kwargs, pairs...)
	})
}

// UnpackArgsIgnoreKeyword 类似于 UnpackArgs，但是忽略命名参数
func (h *Helper) UnpackArgsIgnoreKeyword(pairs ...interface{}) *Helper {
	return h.withErrorHandler(func() error {
		return starlark.UnpackArgs(h.Name, h.PositionalArgs, nil, pairs...)
	})
}

func (h *Helper) UnpackPositionalArgs(min int, vars ...interface{}) *Helper {
	return h.withErrorHandler(func() error {
		return starlark.UnpackPositionalArgs(h.Name, h.PositionalArgs, h.Kwargs, min, vars...)
	})
}

// Deprecated: UnpackBasicArgs  please use UnpackArgs directly
func (h *Helper) UnpackBasicArgs(pairs ...interface{}) *Helper {
	return h.UnpackArgs(pairs...)
}

// GetKeywordArgs 从 keyword args 中提取出指定 key 对应的值，如果值不存在返回 nil
func (h *Helper) GetKeywordArgs(key string) starlark.Value {
	return h.KeywordArgs[key]
}

// GetKeywordArgsString 从 keyword args 中提取出指定 key 对应的值并转换为 string 类型
func (h *Helper) GetKeywordArgsString(key string) (string, error) {
	vv := h.GetKeywordArgs(key)
	if vv == nil {
		return "", ErrArgNotExist{key}
	}

	if v, ok := vv.(starlark.String); ok {
		return string(v), nil
	}

	return "", ErrArgTypeMismatch{key, "string"}
}

// GetKeywordArgsStringWithDefault like GetKeywordArgsString but accept a defaultValue arg
func (h *Helper) GetKeywordArgsStringWithDefault(key string, defaultValue string) (string, error) {
	b, err := h.GetKeywordArgsString(key)
	if IsArgNotExist(err) {
		return defaultValue, nil
	}

	return b, err
}

// GetKeywordArgsBool 从 keyword args 中提取出指定 key 对应的值并转换为 int64 类型
func (h *Helper) GetKeywordArgsBool(key string) (bool, error) {
	vv := h.GetKeywordArgs(key)
	if vv == nil {
		return false, ErrArgNotExist{key}
	}

	if v, ok := vv.(starlark.Bool); ok {
		return bool(v), nil
	}

	return false, ErrArgTypeMismatch{key, "bool"}
}

// GetKeywordArgsBoolWithDefault like GetKeywordArgsBool but accept a defaultValue arg
func (h *Helper) GetKeywordArgsBoolWithDefault(key string, defaultValue bool) (bool, error) {
	b, err := h.GetKeywordArgsBool(key)
	if IsArgNotExist(err) {
		return defaultValue, nil
	}

	return b, err
}

// GetKeywordArgsInt64 从 keyword args 中提取出指定 key 对应的值并转换为 int64 类型
func (h *Helper) GetKeywordArgsInt64(key string) (int64, error) {
	vv := h.GetKeywordArgs(key)
	if vv == nil {
		return 0, ErrArgNotExist{key}
	}

	if v, ok := vv.(starlark.Int); ok {
		if vvv, ok := v.Int64(); ok {
			return vvv, nil
		} else {
			return 0, ErrArgTypeMismatch{key, "int64: overflow"}
		}
	}

	return 0, ErrArgTypeMismatch{key, "int"}
}

// GetKeywordArgsInt64WithDefault like GetKeywordArgsInt64 but accept a defaultValue arg
func (h *Helper) GetKeywordArgsInt64WithDefault(key string, defaultValue int64) (int64, error) {
	b, err := h.GetKeywordArgsInt64(key)
	if IsArgNotExist(err) {
		return defaultValue, nil
	}

	return b, err
}

// GetKeywordArgsUint64 从 keyword args 中提取出指定 key 对应的值并转换为 uint64 类型
func (h *Helper) GetKeywordArgsUint64(key string) (uint64, error) {
	vv := h.GetKeywordArgs(key)
	if vv == nil {
		return 0, ErrArgNotExist{key}
	}

	if v, ok := vv.(starlark.Int); ok {
		if vvv, ok := v.Uint64(); ok {
			return vvv, nil
		} else {
			return 0, ErrArgTypeMismatch{key, "uint64: overflow"}
		}
	}

	return 0, ErrArgTypeMismatch{key, "int (positive)"}
}

// GetKeywordArgsUint64WithDefault like GetKeywordArgsUint64 but accept a defaultValue arg
func (h *Helper) GetKeywordArgsUint64WithDefault(key string, defaultValue uint64) (uint64, error) {
	b, err := h.GetKeywordArgsUint64(key)
	if IsArgNotExist(err) {
		return defaultValue, nil
	}

	return b, err
}

// GetKeywordArgsStringList 从 keyword args 中提取出指定 key 对应的值并转换为 []string 类型
// If type of key is a single string, it's also valid and will return a single element string slice
func (h *Helper) GetKeywordArgsStringList(key string) ([]string, error) {
	vv := h.GetKeywordArgs(key)
	if vv == nil {
		return nil, ErrArgNotExist{key}
	}

	return ToStringList(vv)
}

// GetKeywordArgsStringListWithDefault like GetKeywordArgsStringList but accept a defaultValue arg
func (h *Helper) GetKeywordArgsStringListWithDefault(key string, defaultValue []string) ([]string, error) {
	b, err := h.GetKeywordArgsStringList(key)
	if IsArgNotExist(err) {
		return defaultValue, nil
	}

	return b, err
}

// ArgsCount 返回传入的 args 数量
func (h *Helper) ArgsCount() int {
	return len(h.PositionalArgs) + len(h.Kwargs)
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

	if len(h.PositionalArgs) > 0 {
		return h.PositionalArgs[0], nil
	} else {
		return h.Kwargs[0].Index(1), nil
	}
}

func (h *Helper) GetAllPositionalArgs() []starlark.Value {
	return h.PositionalArgs
}

func (h *Helper) ContainKeywordArg(key string) bool {
	return h.KeywordArgs.Has(key)
}

// CheckContainKeywordArg will return error if there is any keyword arguments passed in not exist.
// If there are other keyword arguments, won't error. If you want to check that, maybe you want to use CheckOnlyContainKeywordArg.
func (h *Helper) CheckContainKeywordArg(keys ...string) error {
	for _, k := range keys {
		if !h.ContainKeywordArg(k) {
			return ErrArgNotExist{k}
		}
	}
	return nil
}

// CheckOnlyContainKeywordArg will return error if there is any keyword arguments except for passed in keys.
// If any of keys not exist, won't error. If you want to check that, maybe you want to use CheckContainKeywordArg
func (h *Helper) CheckOnlyContainKeywordArg(keys ...string) error {
	sets := make(map[string]struct{}, len(keys))
	for _, k := range keys {
		sets[k] = struct{}{}
	}

	for k := range h.KeywordArgs {
		if _, ok := sets[k]; !ok {
			return ErrArgExist{k}
		}
	}

	return nil
}

type ErrArgExist struct{ ArgName string }

func (e ErrArgExist) Error() string {
	return fmt.Sprintf("arg %s: exist", e.ArgName)
}

type ErrArgNotExist struct{ ArgName string }

func (e ErrArgNotExist) Error() string {
	return fmt.Sprintf("arg %s: not exist", e.ArgName)
}

type ErrArgTypeMismatch struct {
	ArgName    string
	ExpectType string
}

func (e ErrArgTypeMismatch) Error() string {
	return fmt.Sprintf("arg %s: invalid type (expect %s)", e.ArgName, e.ExpectType)
}

func IsArgExist(err error) bool {
	return errors.As(err, &ErrArgExist{})
}

func IsArgNotExist(err error) bool {
	return errors.As(err, &ErrArgNotExist{})
}

func IsArgTypeMismatch(err error) bool {
	return errors.As(err, &ErrArgTypeMismatch{})
}
