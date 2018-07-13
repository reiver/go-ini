package initoken

import (
	"fmt"
	"io"
)

type Section struct {
	value string
}

func SomeSection(s string) Section {
	return Section{value: s}
}

func (receiver Section) INIToken() Type {
	return receiver
}

func (receiver Section) INITokenSome() SomeType {
	return receiver
}

func (Section) INITokenSection() {
	// Nothing here.
}

func (receiver Section) String() string {
	return fmt.Sprintf("initoken.SomeSection(%q)", receiver.value)
}

func (receiver Section) Unwrap() string {
	return receiver.value
}

func (receiver Section) WriteTo(writer io.Writer) (int64, error) {
	n, err := io.WriteString(writer, receiver.value)
	n64 := int64(n)

	return n64, err
}
