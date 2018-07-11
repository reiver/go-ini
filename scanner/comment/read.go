package iniscanner_comment

import (
	"github.com/reiver/go-ini/token"

	"bytes"
	"fmt"
	"io"
)

func Read(runeScanner io.RuneScanner) (initoken.Comment, int, error) {
	if nil == runeScanner  {
		return initoken.Comment{}, 0, errNilRuneScanner
	}

	var buffer bytes.Buffer

	var n int

	var notFirst bool

	for {
		r, n2, err := runeScanner.ReadRune()
		n += n2
		if nil != err && io.EOF != err {
			return initoken.SomeComment( buffer.String() ), n, internalInternalErrorComplainer{
				value:    buffer.String(),
				activity: "trying to read rune",
				err:      err,
			}
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
			switch r {
			case ';','#':
				// Nothing here.
			default:
				return initoken.Comment{}, n, internalSyntaxErrorComplainer{
					value:  string(r),
					reason: fmt.Sprintf("not a comment, comments begin with a \"#\" or a \";\" charcter."),
				}
			}
		}

		switch r {
		case '\u000A', // line feed
		     '\u000D', // carriage return
		     '\u000B', // vertical tab
		     '\u0085', // next line
		     '\u2028', // line separator
		     '\u2029': // paragraph separator

			if err := runeScanner.UnreadRune(); nil != err {
				return initoken.SomeComment( buffer.String() ), n, internalInternalErrorComplainer{
					value:     buffer.String(),
					activity: "trying to unread rune",
					err:       err,
				}
			}

			n -= n2

			return initoken.SomeComment( buffer.String() ), n, nil
		}


		buffer.WriteRune(r)
	}
}
