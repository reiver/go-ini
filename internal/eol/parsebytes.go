package inieol

import (
	"unicode/utf8"

	"sourcecode.social/reiver/go-eol/cr"
	"sourcecode.social/reiver/go-eol/lf"
	"sourcecode.social/reiver/go-eol/ls"
	"sourcecode.social/reiver/go-eol/nel"
)

func ParseBytes(bytes []byte) (size int, err error) {
	if len(bytes) <= 0 {
		return 0, nil
	}

	var p []byte = bytes

	var r rune
	var runeSize int
	{
		r, runeSize = utf8.DecodeRune(p)
		if utf8.RuneError == r {
			switch runeSize {
			case 1:
				return 0, errRuneError
			default:
				return 0, errInternalError
			}
		}
	}

	{
		switch r {
		case lf.Rune, cr.Rune, nel.Rune, ls.Rune:
			size += runeSize
			p = p[runeSize:]
		default:
			return 0, nil
		}
	}

	{
		if len(p) <= 0 {
			return size, nil
		}

		switch r {
		case nel.Rune, ls.Rune:
			return size, nil
		}
	}

	{
		switch r {
		case lf.Rune:
			r2, runeSize2 := utf8.DecodeRune(p)
			if utf8.RuneError == r2 {
				switch runeSize2 {
				case 1:
					return 0, errRuneError
				default:
					return 0, errInternalError
				}
			}

			if cr.Rune == r2 {
				size += runeSize2
				p = p[runeSize2:]
			}

		case cr.Rune:
			r2, runeSize2 := utf8.DecodeRune(p)
			if utf8.RuneError == r2 {
				switch runeSize2 {
				case 1:
					return 0, errRuneError
				default:
					return 0, errInternalError
				}
			}

			if lf.Rune == r2 {
				size += runeSize2
				p = p[runeSize2:]
			}
		}
	}

	return size, nil
}
