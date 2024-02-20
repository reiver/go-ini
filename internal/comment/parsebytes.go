package inicomment

import (
	"unicode/utf8"

	"sourcecode.social/reiver/go-eol/cr"
	"sourcecode.social/reiver/go-eol/lf"
	"sourcecode.social/reiver/go-eol/ls"
	"sourcecode.social/reiver/go-eol/nel"
	"sourcecode.social/reiver/go-erorr"

	"github.com/reiver/go-ini/internal/eol"
)

func ParseBytes(bytes []byte) (comment string, size int, err error) {

	if len(bytes) <= 0 {
		return "", 0, nil
	}

	var p []byte = bytes

	var magicSize int
	{
		r, runeSize := utf8.DecodeRune(p)
		if utf8.RuneError == r {
			switch runeSize {
			case 1:
				return "", 0, errRuneError
			default:
				return "", 0, errInternalError
			}
		}

		if !IsMagic(r) {
			return "", 0, erorr.Errorf("ini: not a comment — a comment must begin with a '#' (U+0023) or ';' (U+003B) character")
		}

		magicSize = runeSize
		size += runeSize
		p = p[runeSize:]
	}

	{
		loop: for {
			length := len(p)
			if length <= 0 {
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
				size += runeSize
				p = p[runeSize:]
			}
		}

		comment = string(bytes[magicSize:size])
	}

	{
		eolSize, err := inieol.ParseBytes(p)
		if nil != err {
			return "", 0, err
		}

		size += eolSize
	}

	return comment, size, nil
}
