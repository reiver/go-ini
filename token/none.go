package initoken

type NoneType interface {
	Type

	INITokenNone()
}

type internalNone struct {}

func None() NoneType {
	return internalNone{}
}

func (receiver internalNone) INIToken() Type {
	return receiver
}

func (internalNone) INITokenNone() {
	// Nothing here.
}
