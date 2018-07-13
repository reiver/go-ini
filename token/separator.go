package initoken

import (
	"fmt"
	"io"
)

type Separator struct {
	value string
}

func SomeSeparator(s string) Separator {
	return Separator{value: s}
}

func (receiver Separator) INIToken() Type {
	return receiver
}

func (receiver Separator) INITokenSome() SomeType {
	return receiver
}

func (Separator) INITokenSeparator() {
	// Nothing here.
}

func (receiver Separator) String() string {
	return fmt.Sprintf("initoken.SomeSeparator(%q)", receiver.value)
}

func (receiver Separator) Unwrap() string {
	return receiver.value
}

func (receiver Separator) WriteTo(writer io.Writer) (int64, error) {
	n, err := io.WriteString(writer, receiver.value)
	n64 := int64(n)

	return n64, err
}
