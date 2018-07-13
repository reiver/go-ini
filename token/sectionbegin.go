package initoken

import (
	"fmt"
	"io"
)

type SectionBegin struct {
	value string
}

func SomeSectionBegin(s string) SectionBegin {
	return SectionBegin{value: s}
}

func (receiver SectionBegin) INIToken() Type {
	return receiver
}

func (receiver SectionBegin) INITokenSome() SomeType {
	return receiver
}

func (SectionBegin) INITokenSectionBegin() {
	// Nothing here.
}

func (receiver SectionBegin) String() string {
	return fmt.Sprintf("initoken.SomeSectionBegin(%q)", receiver.value)
}

func (receiver SectionBegin) Unwrap() string {
	return receiver.value
}

func (receiver SectionBegin) WriteTo(writer io.Writer) (int64, error) {
	n, err := io.WriteString(writer, receiver.value)
	n64 := int64(n)

	return n64, err
}
