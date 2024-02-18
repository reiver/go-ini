package inisection

import (
	"github.com/reiver/go-ini/internal/scanner/error"
	"github.com/reiver/go-ini/token"

	"bytes"
	"io"
)


func Read(runeScanner io.RuneScanner) (initoken.Section, int, error) {
	if nil == runeScanner  {
		return initoken.Section{}, 0, iniscanner_error.NilRuneScanner
	}

	var buffer bytes.Buffer

	var n int

	var notFirst bool

	for {
		r, n2, err := runeScanner.ReadRune()
		n += n2
		if nil != err && io.EOF != err {
			return initoken.SomeSection( buffer.String() ), n, iniscanner_error.InternalError(
				buffer.String(),
				"trying to read rune",
				err,
			)
		}
		if io.EOF == err {
			return initoken.Section{}, n, iniscanner_error.SyntaxError(
				"not a section, sections end with a \"]\" charcter",
				string(r),
			)
		}

		if !notFirst {
			notFirst = true
			switch r {
			case '[':
				// Nothing here.
			default:
				return initoken.Section{}, n, iniscanner_error.SyntaxError(
					"not a section, sections begin with a \"[\" charcter",
					string(r),
				)
			}
		}

		if ']' == r {
			buffer.WriteRune(r)

			return initoken.SomeSection( buffer.String() ), n, nil
		}

		buffer.WriteRune(r)
	}
}
