package starlarkhelper

type BasicType string

const (
	TypeNone   BasicType = "NoneType"
	TypeInt    BasicType = "int"
	TypeFloat  BasicType = "float"
	TypeBool   BasicType = "bool"
	TypeString BasicType = "string"
	TypeBytes  BasicType = "bytes"

	TypeAny BasicType = "any"
)
