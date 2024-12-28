package ininame

import (
	"unicode"
	"unicode/utf8"

	"github.com/reiver/go-eol/cr"
	"github.com/reiver/go-eol/lf"
	"github.com/reiver/go-eol/ls"
	"github.com/reiver/go-eol/nel"
)

// ParseBytes looks for an INI 'name' (for an INI name-value pair) at the beginning for 'bytes'.
func ParseBytes(bytes []byte) (name string, size int, err error) {
	if len(bytes) <= 0 {
		return "", 0, nil
	}

	{
		var p []byte = bytes

		var nameBuffer [bufferSize]byte
		var nameBytes []byte = nameBuffer[0:0]

		nameLoop: for {
			length := len(p)
			if length <= 0 {
				name = string(nameBytes)
				break nameLoop
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
			case '\t', ' ', ':', '=':
				break nameLoop
			case lf.Rune, cr.Rune, nel.Rune, ls.Rune:
				break nameLoop
			default:
				var runeLowerCased rune = unicode.ToLower(r)   // ex: 'E' -> 'e'
				var runeAsUTF8 string = string(runeLowerCased) // ex: 'e' -> "e"

				nameBytes = append(nameBytes, runeAsUTF8...)
				p = p[runeSize:]
				size += runeSize
			}
		}

		name = string(nameBytes)
	}

	return name, size, nil
}
