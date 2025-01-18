package ini

import (
	"fmt"
)
// InvalidDestinationTypeError is returned by [Unmarshal] if it is given a 'destination' to unmarshal whose type is not supported by [Unmarshal].
type InvalidDestinationTypeError struct {
	invalidType string
}

func CreateInvalidDestinationTypeError(value any) error {
	return InvalidDestinationTypeError{
		invalidType: fmt.Sprintf("%T", value),
	}
}

func (receiver InvalidDestinationTypeError) Error() string {
	return fmt.Sprintf("ini: cannot unmarshal into something of type %s", receiver.invalidType)
}

func (receiver InvalidDestinationTypeError) InvalidType() string {
	return receiver.invalidType
}
