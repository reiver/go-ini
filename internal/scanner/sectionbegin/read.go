package iniscanner_sectionbegin

import (
	"github.com/reiver/go-ini/internal/scanner/error"
	"github.com/reiver/go-ini/token"

	"bytes"
	"io"
)

func Read(runeScanner io.RuneScanner) (initoken.SectionBegin, int, error) {
	if nil == runeScanner  {
		return initoken.SectionBegin{}, 0, iniscanner_error.NilRuneScanner
	}

	var buffer bytes.Buffer

	var n int

	{
		r, n2, err := runeScanner.ReadRune()
		n += n2
		if nil != err && io.EOF != err {
			return initoken.SomeSectionBegin( buffer.String() ), n, iniscanner_error.InternalError(
				buffer.String(),
				"trying to read rune",
				err,
			)
		}
		if io.EOF == err {
			return initoken.SectionBegin{}, n, iniscanner_error.SyntaxError(
					"not a sectionbegin, sectionbegins are a \"{\" charcter",
					string(r),
				)
		}

		switch {
		case Peek(r):
			// Nothing here.
		default:
			return initoken.SectionBegin{}, n, iniscanner_error.SyntaxError(
				"not a sectionbegin, sectionbegins are a \"{\" charcter",
				string(r),
			)
		}

		buffer.WriteRune(r)
	}

	return initoken.SomeSectionBegin( buffer.String() ), n, nil
}
