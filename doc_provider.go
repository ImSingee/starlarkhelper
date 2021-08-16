package starlarkhelper

import (
	"fmt"
	"go.starlark.net/starlark"
)

// DocProvider 标志着一个 starlark 数据类型可以进行 help 执行
type DocProvider interface {
	starlark.Value
	CanHelp
}

type DocObject struct {
	CanHelp
}

var _ DocProvider = DocObject{}

func NewDocObject(h CanHelp) DocObject {
	return DocObject{CanHelp: h}
}

func (d DocObject) String() string        { return "DocProvider" }
func (d DocObject) Type() string          { return "DocProvider" }
func (d DocObject) Freeze()               {}
func (d DocObject) Truth() starlark.Bool  { return false }
func (d DocObject) Hash() (uint32, error) { return 0, fmt.Errorf("un-hashable") }
