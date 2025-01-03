package ini

import (
	"github.com/reiver/go-erorr"
)

// SectionOf returns the ini-section of a type, if it has one.
//
// Some Go built-in types have a ini-section; such as map[string]string
//
// A custom type can also have an ini-section by implementing the [Sectioner] interface.
func SectionOf(v any, nesting ...string) ([]byte, error) {

	var sectioner Sectioner

	switch casted := v.(type) {
	case Sectioner:
		sectioner = casted
	case map[string]string:
		sectioner = internalMapStringString{casted}
	default:
		return nil, erorr.Errorf("ini: type %T does not have a 'section' representation", v)
	}

	{
		var buffer [256]byte
		var p []byte = buffer[0:0]

		return sectioner.AppendINISection(p, nesting...)
	}
}
