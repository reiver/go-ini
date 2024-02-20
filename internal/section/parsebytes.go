package inisection

import (
	"unicode/utf8"

	"sourcecode.social/reiver/go-eol/cr"
	"sourcecode.social/reiver/go-eol/lf"
	"sourcecode.social/reiver/go-eol/ls"
	"sourcecode.social/reiver/go-eol/nel"
)

func ParseString(str string) (string, error) {
	if len(str) <= 0 {
		return "", errEmptyString
	}

	{
		str0 := str[0]

		if '[' != str0 {
			return "", errMagicNotFound
		}
	}

	var result string = str[1:]

	for {
		length := len(result)

		if length <= 0 {
			return "", nil
		}

		r, size := utf8.DecodeLastRuneInString(result)

		switch r {
		case cr.Rune, lf.Rune, ls.Rune, nel.Rune:
			result = result[:length-size]
		case ']':
			result = result[:length-size]
			return result, nil
		default:
			return result, nil
		}
	}
}
