package initoken

import (
	"fmt"
	"io"
)

type SectionEnd struct {
	value string
}

func SomeSectionEnd(s string) SectionEnd {
	return SectionEnd{value: s}
}

func (receiver SectionEnd) INIToken() Type {
	return receiver
}

func (receiver SectionEnd) INITokenSome() SomeType {
	return receiver
}

func (SectionEnd) INITokenSectionEnd() {
	// Nothing here.
}

func (receiver SectionEnd) String() string {
	return fmt.Sprintf("initoken.SomeSectionEnd(%q)", receiver.value)
}

func (receiver SectionEnd) Unwrap() string {
	return receiver.value
}

func (receiver SectionEnd) WriteTo(writer io.Writer) (int64, error) {
	n, err := io.WriteString(writer, receiver.value)
	n64 := int64(n)

	return n64, err
}
