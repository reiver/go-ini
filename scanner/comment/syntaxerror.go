package iniscanner_comment

import (
	"fmt"
)

type SyntaxErrorComplainer interface {
	error
	SyntaxErrorComplainer()
}

type internalSyntaxErrorComplainer struct{
	value  string
	reason string
}

func (complainer internalSyntaxErrorComplainer) Error() string {
	return fmt.Sprintf("Syntax Error parsing key %q because: %s", complainer.value, complainer.reason)
}

func (complainer internalSyntaxErrorComplainer) SyntaxErrorComplainer() {
	// Nothing here.
}
