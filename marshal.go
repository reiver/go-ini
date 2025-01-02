package ini

import (
	"bytes"
)

// Marshal returns the INI encoding of `value`.
//
// See also [ToString] and [Write]
func Marshal(value any) ([]byte, error) {
	var buffer bytes.Buffer

	err := Write(&buffer, value)
	if nil != err {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// NestedMarshal is similar to [Marshal] except nested.
//
// See also [NestedToString] and [NestedWrite]
func NestedMarshal(value any, nesting ...string) ([]byte, error) {
	var buffer bytes.Buffer

	err := NestedWrite(&buffer, value, nesting...)
	if nil != err {
		return nil, err
	}

	return buffer.Bytes(), nil
}
