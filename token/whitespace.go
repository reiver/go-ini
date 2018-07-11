package initoken

import (
	"io"
)

type Whitespace struct {
	value string
}

func SomeWhitespace(s string) Whitespace {
	return Whitespace{value: s}
}

func (receiver Whitespace) INIToken() Type {
	return receiver
}

func (receiver Whitespace) INITokenSome() SomeType {
	return receiver
}

func (Whitespace) INITokenWhitespace() {
	// Nothing here.
}

func (receiver Whitespace) String() string {
	return string(receiver.value)
}

func (receiver Whitespace) WriteTo(writer io.Writer) (int64, error) {
	n, err := io.WriteString(writer, receiver.value)
	n64 := int64(n)

	return n64, err
}
