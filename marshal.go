package ini

import (
	"bytes"
)

// Marshal returns the INI encoding of `value`.
func Marshal(value any) ([]byte, error) {
	var buffer bytes.Buffer

	err := WriteINI(&buffer, value)
	if nil != err {
		return nil, err
	}

	return buffer.Bytes(), nil
}
