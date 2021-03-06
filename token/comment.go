package initoken

import (
	"fmt"
	"io"
)

type Comment struct {
	value string
}

func SomeComment(s string) Comment {
	return Comment{value: s}
}

func (receiver Comment) INIToken() Type {
	return receiver
}

func (receiver Comment) INITokenSome() SomeType {
	return receiver
}

func (Comment) INITokenComment() {
	// Nothing here.
}

func (receiver Comment) String() string {
	return fmt.Sprintf("initoken.SomeComment(%q)", receiver.value)
}

func (receiver Comment) Unwrap() string {
	return receiver.value
}

func (receiver Comment) WriteTo(writer io.Writer) (int64, error) {
	n, err := io.WriteString(writer, receiver.value)
	n64 := int64(n)

	return n64, err
}
