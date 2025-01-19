package iniampersandvalue

import (
	gobytes "bytes"
)

// ParseBytes looks for an INI 'ampersand-value' (for an INI name-value pair) at the beginning for 'bytes'.
func ParseBytes(end string, bytes []byte) (value string, size int, err error) {
	if len(bytes) <= 0 {
		var nada string
		return nada, 0, nil
	}

	var endBytes []byte
	{
		var buffer [256]byte
		endBytes = buffer[0:0]
		endBytes = append(endBytes, end...)
	}

	{
		var index int = gobytes.Index(bytes, endBytes)

		switch {
		case index < 0:
			value = string(bytes)
			size = len(value)
		default:
			value = string(bytes[:index])
			size = len(value) + len(end)
		}
	}

	return value, size, nil
}
