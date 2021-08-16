package starlarkhelper

import (
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

type Module struct {
	Name string

	Short string // module 功能简介

	PreDeclares []*PreDeclaredValue
	Functions   []*Function

	DataTypes []*DataTypeMember
}

type StarlarkModule struct {
	starlarkstruct.Module

	internal *Module
}

func (m *Module) Get() *StarlarkModule {
	members := make(starlark.StringDict, len(m.PreDeclares)+len(m.Functions))
	for _, member := range m.PreDeclares {
		members[member.Name] = member.getForModule(m.Name)
	}
	for _, member := range m.Functions {
		member.moduleName = m.Name
		members[member.FuncName] = member
	}

	return &StarlarkModule{
		Module: starlarkstruct.Module{
			Name:    m.Name,
			Members: members,
		},
		internal: m,
	}
}
