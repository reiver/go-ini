package ini

import (
	"github.com/reiver/go-cast"
)

// ValueOf returns the ini-string value of a type, if it has one.
//
// Many Go built-in types have a ini-string value.
//
// A custom type can also have an ini-string value by implementing the [Valuer] interface.
func ValueOf(v any) (string, error) {

	switch casted := v.(type) {
	case Valuer:
		return casted.INIValue()
	default:
		return cast.String(v)
	}
}
