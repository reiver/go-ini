package initoken

import (
	"fmt"
	"io"
)

type Key struct {
	value string
}

func SomeKey(s string) Key {
	return Key{value: s}
}

func (receiver Key) INIToken() Type {
	return receiver
}

func (receiver Key) INITokenSome() SomeType {
	return receiver
}

func (Key) INITokenKey() {
	// Nothing here.
}

func (receiver Key) String() string {
	return fmt.Sprintf("initoken.SomeKey(%q)", receiver.value)
}

func (receiver Key) Unwrap() string {
	return receiver.value
}

func (receiver Key) WriteTo(writer io.Writer) (int64, error) {
	n, err := io.WriteString(writer, receiver.value)
	n64 := int64(n)

	return n64, err
}
