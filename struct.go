package starlarkhelper

import "go.starlark.net/starlark"

type Struct struct {
	Name string // struct 名称
}

func (s Struct) String() string {
	panic("implement me")
}

func (s Struct) Type() string {
	panic("implement me")
}

func (s Struct) Freeze() {
	panic("implement me")
}

func (s Struct) Truth() starlark.Bool {
	panic("implement me")
}

func (s Struct) Hash() (uint32, error) {
	panic("implement me")
}

func (s Struct) Attr(name string) (starlark.Value, error) {
	panic("implement me")
}

func (s Struct) AttrNames() []string {
	panic("implement me")
}
