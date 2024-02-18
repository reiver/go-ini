package iniscanner_comment

import (
	"github.com/reiver/go-ini/internal/scanner/error"
	"github.com/reiver/go-ini/token"

	"github.com/reiver/go-whitespace"

	"bytes"
	"fmt"
	"io"
)

func Read(runeScanner io.RuneScanner) (initoken.Comment, int, error) {
	if nil == runeScanner  {
		return initoken.Comment{}, 0, iniscanner_error.NilRuneScanner
	}

	var buffer bytes.Buffer

	var n int

	var notFirst bool

	for {
		r, n2, err := runeScanner.ReadRune()
		n += n2
		if nil != err && io.EOF != err {
			return initoken.SomeComment( buffer.String() ), n, iniscanner_error.InternalError(
				buffer.String(),
				"trying to read rune",
				err,
			)
		}
		if io.EOF == err {
			if 0 < n {
				return initoken.SomeComment( buffer.String() ), n, nil
			} else {
				return initoken.SomeComment( buffer.String() ), n, io.EOF
			}
		}

		if !notFirst {
			notFirst = true
			switch {
			case Peek(r):
				// Nothing here.
			default:
				return initoken.Comment{}, n, iniscanner_error.SyntaxError(
					fmt.Sprintf("not a comment, comments begin with a \"#\" or a \";\" charcter"),
					string(r),
				)
			}
		}

		if whitespace.IsMandatoryBreak(r) {

			if err := runeScanner.UnreadRune(); nil != err {
				return initoken.SomeComment( buffer.String() ), n, iniscanner_error.InternalError(
					buffer.String(),
					"trying to unread rune",
					err,
				)
			}

			n -= n2

			return initoken.SomeComment( buffer.String() ), n, nil
		}

		buffer.WriteRune(r)
	}
}
