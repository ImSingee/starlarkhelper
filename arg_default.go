package starlarkhelper

import (
	"fmt"
	"strconv"
)

type defaultArg struct {
	Arg

	defaultValue interface{}
}

func (n *defaultArg) InDefaultValue() string        { return defaultValueToString(n.defaultValue) }
func (n *defaultArg) Default(value interface{}) Arg { n.defaultValue = value; return n }

// withDefault 标识默认类型
func withDefault(arg Arg, defaultValue interface{}) Arg {
	return &defaultArg{Arg: arg, defaultValue: defaultValue}
}

type RawDefaultString string

// defaultValueToString 将默认值转为 string
// Panic: 该函数专为默认值设计，如果传递的参数类型不在 nil, int, float64, bool, string 中会 panic
func defaultValueToString(v interface{}) string {
	switch v := v.(type) {
	case nil:
		return ""
	case RawDefaultString:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case float64:
		return fmt.Sprintf("%v", v)
	case string:
		return strconv.Quote(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		panic("not implemented")
	}
}
