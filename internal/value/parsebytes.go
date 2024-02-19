package inivalue

import (
	"unicode/utf8"

	"sourcecode.social/reiver/go-eol/cr"
	"sourcecode.social/reiver/go-eol/lf"
	"sourcecode.social/reiver/go-eol/ls"
	"sourcecode.social/reiver/go-eol/nel"
)

// ParseBytes looks for an INI 'value' (for an INI name-value pair) at the beginning for 'bytes'.
func ParseBytes(bytes []byte) (value string, size int, err error) {
	if len(bytes) <= 0 {
		return "", 0, nil
	}

	{
		var p []byte = bytes

		var valueBuffer [bufferSize]byte
		var valueBytes []byte = valueBuffer[0:0]

		loop: for {
			length := len(p)
			if length <= 0 {
				value = string(valueBytes)
				break loop
			}

			r, runeSize := utf8.DecodeRune(p)
			if utf8.RuneError == r {
				switch runeSize {
				case 1:
					return "", 0, errRuneError
				default:
					return "", 0, errInternalError
				}
			}

			switch r {
			case lf.Rune, cr.Rune, nel.Rune, ls.Rune:
				break loop
			default:
				var runeAsUTF8 string = string(r) // ex: 'E' -> "E"

				valueBytes = append(valueBytes, runeAsUTF8...)
				p = p[runeSize:]
				size += runeSize
			}
		}

		value = string(valueBytes)
	}

	return value, size, nil
}
