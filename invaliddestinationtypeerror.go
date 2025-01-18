package ini

import (
	"fmt"
)
// InvalidDestinationTypeError is returned by [Unmarshal] if it is given a 'destination' to unmarshal whose type is not supported by [Unmarshal].
//
// You can create an InvalidDestinationTypeError using [CreateInvalidDestinationTypeError].
//
// InvalidDestinationTypeError has a method named [InvalidDestinationTypeError.InvalidType] that will return the name of the invalid-type a Go `string`.
type InvalidDestinationTypeError struct {
	invalidType string
}

// CreateInvalidDestinationTypeError returns a [InvalidDestinationTypeError].
func CreateInvalidDestinationTypeError(value any) error {
	return InvalidDestinationTypeError{
		invalidType: fmt.Sprintf("%T", value),
	}
}

// Error makes [InvalidDestinationTypeError] fit the built-in Go `error` interface.
func (receiver InvalidDestinationTypeError) Error() string {
	return fmt.Sprintf("ini: cannot unmarshal into something of type %s", receiver.invalidType)
}

// InvalidType returns the name of the invalid-type as a Go `string`.
func (receiver InvalidDestinationTypeError) InvalidType() string {
	return receiver.invalidType
}
