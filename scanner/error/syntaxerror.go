package iniscanner_error

import (
	"fmt"
)

type SyntaxErrorComplainer interface {
	error
	SyntaxErrorComplainer()
}

func SyntaxError(reason string, value string) error {
	return internalSyntaxErrorComplainer{
		reason: reason,
		value: value,
	}
}

type internalSyntaxErrorComplainer struct{
	reason string
	value  string
}

func (complainer internalSyntaxErrorComplainer) Error() string {
	return fmt.Sprintf("Syntax Error parsing key %q because: %s", complainer.value, complainer.reason)
}

func (complainer internalSyntaxErrorComplainer) SyntaxErrorComplainer() {
	// Nothing here.
}
