package ini

import (
	"io"
)

// AppendKeyValue appends the INI representation of a key-vaue pair.
//
// You might use this function is you are implementing the [Marshaler] interface for a custom type.
// And in particular, use it (directly or indirectly) from you Marshal method.
//
// Also see [KeyValueToString] and [WriteKeyValue]
func AppendKeyValue(p []byte, key string, value string) []byte {
	p = EncodeAndAppendKey(p, key)
	p = append(p, " = "...)
	p = EncodeAndAppendValue(p, value)
	p = append(p, '\n')

	return p
}

// See also [AppendKeyValue] and [WriteKeyValue]
func KeyValueToString(key string, value string) string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = AppendKeyValue(p, key, value)

	return string(p)
}

// See also [AppendKeyValue] and [KeyValueToString]
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
