package ini

import (
	"io"
)

func AppendKeyValue(p []byte, key string, value string) []byte {
	p = EncodeAndAppendKey(p, key)
	p = append(p, " = "...)
	p = EncodeAndAppendValue(p, value)
	p = append(p, '\n')

	return p
}

func WriteKeyValue(dst io.Writer, key string, value string) error {
	if nil == dst {
		return errNilWriter
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = AppendKeyValue(p, key, value)

	_, err := dst.Write(p)
	if nil != err {
		return err
	}

	return nil
}
