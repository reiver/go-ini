package initoken

import (
	"fmt"
	"io"
)

type Value struct {
	value string
}

func SomeValue(s string) Value {
	return Value{value: s}
}

func (receiver Value) INIToken() Type {
	return receiver
}

func (receiver Value) INITokenSome() SomeType {
	return receiver
}

func (Value) INITokenValue() {
	// Nothing here.
}

func (receiver Value) String() string {
	return fmt.Sprintf("initoken.SomeValue(%q)", receiver.value)
}

func (receiver Value) Unwrap() string {
	return receiver.value
}

func (receiver Value) WriteTo(writer io.Writer) (int64, error) {
	n, err := io.WriteString(writer, receiver.value)
	n64 := int64(n)

	return n64, err
}
