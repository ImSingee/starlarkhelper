package starlarkhelper

import (
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

type Module struct {
	Name string

	Short string // module 功能简介

	SubModules  []*Module
	PreDeclares []*PreDeclaredValue
	Functions   []*Function
	Structs     []*StructDescription

	FuncMiddleware Middleware

	//DataTypes []*DataTypeMember
}

type StarlarkModule struct {
	starlarkstruct.Module

	internal *Module
}

// Get 获取 StarlarkModule
func (m Module) Get() *StarlarkModule {
	members := make(starlark.StringDict, len(m.PreDeclares)+len(m.Functions))
	for _, member := range m.SubModules {
		member = member.Copy()
		member.FuncMiddleware = ChainMiddleware(m.FuncMiddleware, member.FuncMiddleware)
		members[member.Name] = member.Get()
	}
	for _, member := range m.PreDeclares {
		members[member.Name] = member.getForModule(m.Name)
	}
	for _, member := range m.Functions {
		member = member.Copy()
		member.Middleware = ChainMiddleware(m.FuncMiddleware, member.Middleware)
		member.moduleName = m.Name
		members[member.FuncName] = member
	}
	for _, member := range m.Structs {
		member.moduleName = m.Name
		members[member.Name] = member.ProvideDoc()
	}

	return &StarlarkModule{
		Module: starlarkstruct.Module{
			Name:    m.Name,
			Members: members,
		},
		internal: &m,
	}
}

func (m Module) Copy() *Module {
	return &m
}
