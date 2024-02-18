package iniscanner_separator

import (
	"github.com/reiver/go-ini/internal/scanner/error"
	"github.com/reiver/go-ini/token"

	"bytes"
	"io"
)

func Read(runeScanner io.RuneScanner) (initoken.Separator, int, error) {
	if nil == runeScanner  {
		return initoken.Separator{}, 0, iniscanner_error.NilRuneScanner
	}

	var buffer bytes.Buffer

	var n int

	{
		r, n2, err := runeScanner.ReadRune()
		n += n2
		if nil != err && io.EOF != err {
			return initoken.SomeSeparator( buffer.String() ), n, iniscanner_error.InternalError(
				buffer.String(),
				"trying to read rune",
				err,
			)
		}
		if io.EOF == err {
			return initoken.Separator{}, n, iniscanner_error.SyntaxError(
					"not a separator, separators are a \"=\" or a \":\" charcter",
					string(r),
				)
		}

		switch {
		case Peek(r):
			// Nothing here.
		default:
			return initoken.Separator{}, n, iniscanner_error.SyntaxError(
				"not a separator, separators are a \"=\" or a \":\" charcter",
				string(r),
			)
		}

		buffer.WriteRune(r)
	}

	return initoken.SomeSeparator( buffer.String() ), n, nil
}
