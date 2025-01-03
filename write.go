package ini

import (
	"io"
)

// Write writes the INI encoding of `src` to `dst`.
//
// See also [Marshal] and [ToString]
func Write(dst io.Writer, src any) error {
	if nil == dst {
		return errNilWriter
	}

	p, err := Marshal(src)
	if nil != err {
		return err
	}

	_, err = dst.Write(p)
	if nil != err {
		return err
	}

	return nil
}
