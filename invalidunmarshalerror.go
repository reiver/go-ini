package ini

import (
	"fmt"
)

type InvalidUnmarshalError struct {
	typ string
}

func invalidUnmarshalError(value any) error {
	return InvalidUnmarshalError{
		typ: fmt.Sprintf("%T", value),
	}
}

func (receiver InvalidUnmarshalError) Error() string {
	return fmt.Sprintf("ini: cannot unmarshal into %s", receiver.typ)
}
