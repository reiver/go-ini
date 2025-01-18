package ini

import (
	"fmt"
)

// InvalidUnmarshalError is returned by [Unmarshal] if it is given a 'destination' to unmarshal whose type is not supported by [Unmarshal].
type InvalidUnmarshalError struct {
	invalidType string
}

func invalidUnmarshalError(value any) error {
	return InvalidUnmarshalError{
		invalidType: fmt.Sprintf("%T", value),
	}
}

func (receiver InvalidUnmarshalError) Error() string {
	return fmt.Sprintf("ini: cannot unmarshal into something of type %s", receiver.invalidType)
}

func (receiver InvalidUnmarshalError) InvalidType() string {
	return receiver.invalidType
}
