package iniscanner_whitespace

import (
	"github.com/reiver/go-ini/internal/scanner/error"
	"github.com/reiver/go-ini/token"

	"github.com/reiver/go-whitespace"

	"bytes"
	"fmt"
	"io"
)


func Read(runeScanner io.RuneScanner) (initoken.Whitespace, int, error) {
	if nil == runeScanner  {
		return initoken.Whitespace{}, 0, iniscanner_error.NilRuneScanner
	}

	var buffer bytes.Buffer

	var n int

	for {
		r, n2, err := runeScanner.ReadRune()
		n += n2
		if nil != err && io.EOF != err {
//@TODO: Wrap error?
			return initoken.SomeWhitespace( buffer.String() ), n, err
		}
		if io.EOF == err {
			if 0 < n {
				return initoken.SomeWhitespace( buffer.String() ), n, nil
			} else {
				return initoken.SomeWhitespace( buffer.String() ), n, io.EOF
			}
		}

		if !whitespace.IsWhitespace(r) {
			if err := runeScanner.UnreadRune(); nil != err {
//@TODO: Wrap error?
				return initoken.SomeWhitespace( buffer.String() ), n, err
			}

			n -= n2

			if 0 < buffer.Len() {
				return initoken.SomeWhitespace( buffer.String() ), n, nil
			} else {
				return initoken.Whitespace{}, n, iniscanner_error.SyntaxError(
					fmt.Sprintf("not whitespace, whitespace does not begin with a %q charcter", r),
					string(r),
				)
			}
		}

		buffer.WriteRune(r)
	}
}
