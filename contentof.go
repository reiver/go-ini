package ini

import (
	"github.com/reiver/go-erorr"
)

// ContentOf returns the ini-content of a type, if it has one.
//
// Some Go built-in types have an ini-content represention; such as map[string]any
//
// A custom type can also have an ini-content by implementing the [Contenter] interface.
func ContentOf(v any, nesting ...string) ([]byte, error) {

	switch casted := v.(type) {
	case Contenter:
		return casted.INIContent(nesting...)
	case map[string]any:
		return contentOfMap(casted, nesting...)
	default:
		return nil, erorr.Errorf("ini: type %T does not have a 'content' representation", v)
	}
}
