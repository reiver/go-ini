package initoken

import (
	"fmt"
	"io"
)

type Undefined struct {
	value string
}

func SomeUndefined(s string) Undefined {
	return Undefined{value: s}
}

func (receiver Undefined) INIToken() Type {
	return receiver
}

func (receiver Undefined) INITokenSome() SomeType {
	return receiver
}

func (receiver Undefined) String() string {
	return fmt.Sprintf("initoken.SomeUndefined(%q)", receiver.value)
}

func (receiver Undefined) Unwrap() string {
	return receiver.value
}

func (receiver Undefined) WriteTo(writer io.Writer) (int64, error) {
	n, err := io.WriteString(writer, receiver.value)
	n64 := int64(n)

	return n64, err
}
