package ininamevalue

import (
	"unicode/utf8"

	"sourcecode.social/reiver/go-eol/cr"
	"sourcecode.social/reiver/go-eol/lf"
	"sourcecode.social/reiver/go-eol/ls"
	"sourcecode.social/reiver/go-eol/nel"
)

func ParseString(str string) (name string, value string, err error) {
	if len(str) <= 0 {
		return "", "", errEmptyString
	}

	var s string = str

	{
		var nameBuffer [bufferSize]byte
		var nameBytes []byte = nameBuffer[0:0]
		nameLoop: for {
			length := len(s)
			if length <= 0 {
				return string(nameBytes), "", nil
			}

			r, size := utf8.DecodeRuneInString(s)
			if utf8.RuneError == r {
				switch size {
				case 1:
					return "", "", errRuneError
				default:
					return "", "", errInternalError
				}
			}

			switch r {
			case '\t', ' ', ':', '=':
				break nameLoop
			case lf.Rune, cr.Rune, nel.Rune, ls.Rune:
				break nameLoop
			default:
				s = s[size:]
				nameBytes = append(nameBytes, string(r)...)
			}
		}

		name = string(nameBytes)
	}

	{
		gapLoop1: for {
			length := len(s)
			if length <= 0 {
				return name, "", nil
			}

			r, size := utf8.DecodeRuneInString(s)
			if utf8.RuneError == r {
				switch size {
				case 1:
					return name, "", errRuneError
				default:
					return name, "", errInternalError
				}
			}

			switch r {
			case '\t', ' ':
				s = s[size:]
			default:
				break gapLoop1
			}
		}
	}

	{
		r, size := utf8.DecodeRuneInString(s)
		if utf8.RuneError == r {
			switch size {
			case 1:
				return name, "", errRuneError
			default:
				return name, "", errInternalError
			}
		}

		switch r {
		case ':', '=':
			s = s[size:]
		}
	}

	{
		gapLoop2: for {
			length := len(s)
			if length <= 0 {
				return name, "", nil
			}

			r, size := utf8.DecodeRuneInString(s)
			if utf8.RuneError == r {
				switch size {
				case 1:
					return name, "", errRuneError
				default:
					return name, "", errInternalError
				}
			}

			switch r {
			case '\t', ' ':
				s = s[size:]
			default:
				break gapLoop2
			}
		}
	}

	{
		r, size := utf8.DecodeLastRuneInString(s)

		switch r {
		case lf.Rune:
			s = s[:len(s)-size]
			{
				r2, size2 := utf8.DecodeLastRuneInString(s)
				if cr.Rune == r2 {
					s = s[:len(s)-size2]
				}
			}
		case cr.Rune:
			s = s[:len(s)-size]
			{
				r2, size2 := utf8.DecodeLastRuneInString(s)
				if lf.Rune == r2 {
					s = s[:len(s)-size2]
				}
			}
		case nel.Rune:
			s = s[:len(s)-size]
		case ls.Rune:
			s = s[:len(s)-size]
		default:
			// Nothing here.
		}

		value = s
	}

	return name, value, nil
}
