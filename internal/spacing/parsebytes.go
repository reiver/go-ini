package inispacing

import (
	"unicode/utf8"
)

func ParseBytes(bytes []byte) (size int, err error) {

	var p []byte = bytes

	{
		loop: for {
			length := len(p)
			if length <= 0 {
				return size, nil
			}

			r, runeSize := utf8.DecodeRune(p)
			if utf8.RuneError == r {
				switch runeSize {
				case 1:
					return 0, errRuneError
				default:
					return 0, errInternalError
				}
			}

			switch r {
			case '\t', ' ':
				size += runeSize
				p = p[runeSize:]
			default:
				break loop
			}
		}
	}

	return size, nil
}
