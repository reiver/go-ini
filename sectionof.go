package ini

import (
	"github.com/reiver/go-erorr"
)

// SectionOf returns the ini-section of a type, if it has one.
//
// Some Go built-in types have a ini-section; such as map[string]string
//
// A custom type can also have an ini-section by implementing the [Sectioner] interface.
func SectionOf(v any) ([]byte, error) {

	switch casted := v.(type) {
	case Sectioner:
		return casted.INISection()
	case map[string]string:
		return sectionOfMap(casted)
	default:
		return nil, erorr.Errorf("ini: type %T does not have a 'section' representation", v)
	}
}
