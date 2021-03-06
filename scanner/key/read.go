package iniscanner_key

import (
	"github.com/reiver/go-ini/scanner/error"
	"github.com/reiver/go-ini/token"

	"github.com/reiver/go-whitespace"

	"bytes"
	"fmt"
	"io"
)


func Read(runeScanner io.RuneScanner) (initoken.Key, int, error) {
	if nil == runeScanner  {
		return initoken.Key{}, 0, iniscanner_error.NilRuneScanner
	}

	var buffer bytes.Buffer

	var n int

	var notFirst bool

	var escaped bool

	for {
		r, n2, err := runeScanner.ReadRune()
		n += n2
		if nil != err && io.EOF != err {
			return initoken.SomeKey( buffer.String() ), n, iniscanner_error.InternalError(
				buffer.String(),
				"trying to read rune",
				err,
			)
		}
		if io.EOF == err {
			return initoken.Key{}, n, iniscanner_error.SyntaxError(
				"EOF in the middle of a key",
				buffer.String(),
			)
		}

		if !notFirst {
			notFirst = true
			switch r {
			case '=',':',';','#':
				return initoken.Key{}, n, iniscanner_error.SyntaxError(
					fmt.Sprintf("not a key, keys do not begin with a %q or a charcter", r),
					string(r),
				)
			default:
				// Nothing here.
			}
		}


		if '=' == r || ':' == r || ';' == r || '#' == r || whitespace.IsWhitespace(r) {
			if escaped {
				escaped = false
				buffer.WriteRune(r)
			} else {
				if err := runeScanner.UnreadRune(); nil != err {
					return initoken.SomeKey( buffer.String() ), n, iniscanner_error.InternalError(
						buffer.String(),
						"trying to unread rune",
						err,
					)
				}

				n -= n2

				return initoken.SomeKey( buffer.String() ), n, nil
			}
		} else if '\\' == r {
			if escaped {
				escaped = false
				buffer.WriteRune(r)
			} else {
				escaped = true
			}
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
				return initoken.Key{}, n, iniscanner_error.SyntaxError(
					"unknown escape sequence",
					buffer.String(),
				)
			}
		} else {
			buffer.WriteRune(r)
		}
	}
}
