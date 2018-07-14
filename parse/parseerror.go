package iniparse

import (
	"fmt"
)

type ParseError struct {
	value  string
	reason string
}

func (receiver ParseError) Error() string {
	return fmt.Sprintf("ini: parse error of %q: %s", receiver.value, receiver.reason)
}
