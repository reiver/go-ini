package iniscanner_value

import (
	"github.com/reiver/go-ini/scanner/error"
	"github.com/reiver/go-ini/token"

	"github.com/reiver/go-whitespace"

	"bytes"
	"io"
)


func Read(runeScanner io.RuneScanner) (initoken.Value, int, error) {
	if nil == runeScanner  {
		return initoken.Value{}, 0, iniscanner_error.NilRuneScanner
	}

	var buffer bytes.Buffer

	var n int

	var escaped bool

	for {
		r, n2, err := runeScanner.ReadRune()
		n += n2
		if nil != err && io.EOF != err {
			return initoken.SomeValue( buffer.String() ), n, iniscanner_error.InternalError(
				buffer.String(),
				"trying to read rune",
				err,
			)
		}
		if io.EOF == err {
			return initoken.SomeValue( buffer.String() ), n, nil
		}

		if '\\' == r && !escaped {
			escaped = true

		} else if '\\' == r && escaped {
			escaped = false
			buffer.WriteRune(r)

		} else if escaped {
			escaped = false
			switch r {
			case '0':
				buffer.WriteRune(0)
			case 'a':
				buffer.WriteRune('\a')
			case 'b':
				buffer.WriteRune('\b')
			case 't':
				buffer.WriteRune('\t')
			case 'r':
				buffer.WriteRune('\r')
			case 'n':
				buffer.WriteRune('\n')
			case ';':
				buffer.WriteRune(';')
			case '#':
				buffer.WriteRune('#')
			case '=':
				buffer.WriteRune('=')
			case ':':
				buffer.WriteRune(':')
			default:
				switch {
				case whitespace.IsMandatoryBreak(r):
					buffer.WriteRune(r)
				default:
					return initoken.Value{}, n, iniscanner_error.SyntaxError(
						"unknown escape sequence",
						buffer.String(),
					)
				}
			}

		} else if whitespace.IsMandatoryBreak(r) {
			if err := runeScanner.UnreadRune(); nil != err {
				return initoken.SomeValue( buffer.String() ), n, iniscanner_error.InternalError(
					buffer.String(),
					"trying to unread rune",
					err,
				)
			}

			n -= n2

			return initoken.SomeValue( buffer.String() ), n, nil

		} else {
			buffer.WriteRune(r)

		}
	}
}
