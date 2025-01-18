package inisequence

import (
	"unicode/utf8"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-ini/internal/section"
)

func ParseBytes(bytes []byte) (name []string, size int, err error) {

	var nada []string

	if nil == bytes {
		return nada, 0, errNilBytes
	}

	{
		const minlen int = magicLen + endMagicLen

		if actualLen := len(bytes); actualLen < (Size1stMagic + Size2ndMagic) {
			return nada, 0, erorr.Errorf("ini: an INI 'sequence' must be at least %d bytes long — only was %d bytes long (%q)", minlen, actualLen, bytes)
		}
	}

	var size1 int
	{
		var r1 rune
		r1, size1 = utf8.DecodeRune(bytes)
		if utf8.RuneError == r1 {
			switch size {
			case 1:
				return nada, 0, errRuneError
			default:
				return nada, 0, errInternalError
			}
		}

		if !Is1stMagic(r1) {
			return nada, 0, erorr.Errorf("ini: the first character of an INI 'sequence' must be a %q (%U); instead got %q (%U)", magic1stRune, magic1stRune, r1, r1)
		}
	}

	{
		var r2 rune
		r2, size = utf8.DecodeRune(bytes[size1:])
		if utf8.RuneError == r2 {
			switch size {
			case 1:
				return nada, 0, errRuneError
			default:
				return nada, 0, errInternalError
			}
		}

		if !Is2ndMagic(r2) {
			return nada, 0, erorr.Errorf("ini: the second character of an INI 'sequence' must be a %q (%U); instead got %q (%U)", magic2ndRune, magic2ndRune, r2, r2)
		}
	}

	name, size, err = inisection.ParseBytes(bytes[Size1stMagic:])

	return name, Size1stMagic+size2ndEndMagic+size, err

}
