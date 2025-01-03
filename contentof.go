package ini

import (
	"github.com/reiver/go-erorr"
)

// ContentOf returns the ini-content of a type, if it has one.
//
// Some Go built-in types have an ini-content represention; such as map[string]any and map[string]string
//
// A custom type can also have an ini-content by implementing the [Marshaler] interface.
func ContentOf(v any, nesting ...string) ([]byte, error) {

	var marshaler Marshaler

	switch casted := v.(type) {
	case Marshaler:
		marshaler = casted
	case map[string]string:
		marshaler = internalMapStringString{casted}
	case map[string]any:
		marshaler = internalMapStringAny{casted}
	default:
		return nil, erorr.Errorf("ini: type %T does not have a 'content' representation", v)
	}

	{
		var buffer [256]byte
		var p []byte = buffer[0:0]

		return marshaler.MarshalINI(p, nesting...)
	}
}
