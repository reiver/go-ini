package inisection

import (
	"unicode/utf8"

	"github.com/reiver/go-eol/cr"
	"github.com/reiver/go-eol/lf"
	"github.com/reiver/go-eol/ls"
	"github.com/reiver/go-eol/nel"
	"github.com/reiver/go-erorr"

	"github.com/reiver/go-ini/internal/eol"
)

func ParseBytes(bytes []byte) (section string, size int, err error) {

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
			return "", 0, erorr.Errorf("ini: not a section — a seciton must begin with a '[' (U+005B) character")
		}

		magicSize = runeSize
		size += runeSize
		p = p[runeSize:]
	}

	{
		var r rune
		var runeSize int

		loop: for {
			length := len(p)
			if length <= 0 {
				break loop
			}

			r, runeSize = utf8.DecodeRune(p)
			if utf8.RuneError == r {
				switch runeSize {
				case 1:
					return "", 0, errRuneError
				default:
					return "", 0, errInternalError
				}
			}

			switch r {
			case ']':
				break loop
			case lf.Rune, cr.Rune, nel.Rune, ls.Rune:
				break loop
			default:
				size += runeSize
				p = p[runeSize:]
			}
		}

		section = string(bytes[magicSize:size])

		if ']' == r {
			size += runeSize
			p = p[runeSize:]
		}
	}

	{
		eolSize, err := inieol.ParseBytes(p)
		if nil != err {
			return "", 0, err
		}

		size += eolSize
	}

	return section, size, nil
}
