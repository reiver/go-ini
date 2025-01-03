package ini

import (
	"github.com/reiver/go-erorr"
)

// Marshal returns the INI encoding of `content`.
//
// Some Go built-in types have an ini-content represention; such as map[string]any and map[string]string
//
// A custom type can also have an ini-content by implementing the [Marshaler] interface.
//
// See also [ToString] and [Write]
func Marshal(content any, nesting ...string) ([]byte, error) {
	var marshaler Marshaler

	switch casted := content.(type) {
	case Marshaler:
		marshaler = casted
	case map[string]string:
		marshaler = internalMapStringString{casted}
	case map[string]any:
		marshaler = internalMapStringAny{casted}
	default:
		return nil, erorr.Errorf("ini: type %T does not have a 'content' representation", content)
	}

	{
		var buffer [256]byte
		var p []byte = buffer[0:0]

		return marshaler.MarshalINI(p, nesting...)
	}
}
