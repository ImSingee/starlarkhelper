package starlarkhelper

import (
	"fmt"

	"go.starlark.net/starlark"
)

func CheckArgsContain(kwargs []starlark.Tuple, key ...string) bool {
	for _, k := range key {
		for _, pair := range kwargs {
			kk := pair[0].(starlark.String).GoString()

			if kk == k {
				return true
			}
		}
	}

	return false
}

// GetFromArgs 从传入的 keyword args 中提取出指定 key 对应的值
// kwargs 必须是函数所接收到的，否则会 panic
// 如果 key 不存在返回 nil
func GetFromArgs(kwargs []starlark.Tuple, key string) starlark.Value {
	for _, pair := range kwargs {
		kk := pair[0].(starlark.String).GoString()

		if kk == key {
			return pair[1]
		}
	}

	return nil
}

// GetFromArgsInt64 从传入的 keyword args 中提取出指定 key 对应的值
// kwargs 必须是函数所接收到的，否则会 panic
func GetFromArgsInt64(kwargs []starlark.Tuple, key string, defaultValue int64) (int64, error) {
	vv := GetFromArgs(kwargs, key)
	if vv == nil {
		return defaultValue, nil
	}

	if v, ok := vv.(starlark.Int); ok {
		if vvv, ok := v.Int64(); ok {
			return vvv, nil
		} else {
			return 0, fmt.Errorf("invalid arg: int too big")
		}
	}

	return 0, fmt.Errorf("invalid arg type")
}

// GetFromArgsBool 从传入的 keyword args 中提取出指定 key 对应的值
// kwargs 必须是函数所接收到的，否则会 panic
func GetFromArgsBool(kwargs []starlark.Tuple, key string, defaultValue bool) (bool, error) {
	vv := GetFromArgs(kwargs, key)
	if vv == nil {
		return defaultValue, nil
	}

	if v, ok := vv.(starlark.Bool); ok {
		return bool(v), nil
	}

	return false, fmt.Errorf("invalid arg type")
}
