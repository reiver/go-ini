package inicomment

import (
	"unicode/utf8"

	"sourcecode.social/reiver/go-eol/cr"
	"sourcecode.social/reiver/go-eol/lf"
	"sourcecode.social/reiver/go-eol/ls"
	"sourcecode.social/reiver/go-eol/nel"
	"sourcecode.social/reiver/go-erorr"

	"github.com/reiver/go-ini/internal/eol"
)

func ParseBytes(bytes []byte) (comment string, size int, err error) {

	if len(bytes) <= 0 {
		return "", 0, nil
	}

	var p []byte = bytes

	var magicSize int
	{
		r, runeSize := utf8.DecodeRune(p)
		if utf8.RuneError == r {
			switch runeSize {
			case 1:
				return "", 0, errRuneError
			default:
				return "", 0, errInternalError
			}
		}

		if !IsMagic(r) {
			return "", 0, erorr.Errorf("ini: not a comment — a comment must begin with a '#' (U+0023) or ';' (U+003B) character")
		}

		magicSize = runeSize
		size += runeSize
		p = p[runeSize:]
	}

	{
		loop: for {
			length := len(p)
			if length <= 0 {
				break loop
			}

			r, runeSize := utf8.DecodeRune(p)
			if utf8.RuneError == r {
				switch runeSize {
				case 1:
					return "", 0, errRuneError
				default:
					return "", 0, errInternalError
				}
			}

			switch r {
			case lf.Rune, cr.Rune, nel.Rune, ls.Rune:
				break loop
			default:
				size += runeSize
				p = p[runeSize:]
			}
		}

		comment = string(bytes[magicSize:size])
	}

	{
		eolSize, err := inieol.ParseBytes(p)
		if nil != err {
			return "", 0, err
		}

		size += eolSize
	}

	return comment, size, nil
}


/*
func Copy(runeWriter runewriter.RuneWriter, runeScanner io.RuneScanner) error {
	if nil == runeScanner  {
		return errNilRuneScanner
	}
	if nil == runeWriter {
		return errNilRuneWriter
	}

	// handle the first character differently, since it must be a '#' or ';' character.
	{
		r, err := readrune(runeScanner)
		if io.EOF == err {
			return erorr.Errorf("ini: not a comment — a comment must at least start with a '#' (U+0023) or ';' (U+003B) character")
		}
		if nil != err {
			return err
		}
		if !IsMagic(r) {
			return erorr.Errorf("ini: not a comment — a comment does not begin with a %q (%U)", r, r)
		}

		{
			_, err := runeWriter.WriteRune(r)
			if nil != err {
				return erorr.Errorf("ini: problem writing rune: %w", err)
			}
		}
	}

	for {
		r, err := readrune(runeScanner)
		if io.EOF == err {
			return nil
		}
		if nil != err {
			return err
		}

		{
			_, err := runeWriter.WriteRune(r)
			if nil != err {
				return erorr.Errorf("ini: problem writing rune: %w", err)
			}
		}

		switch r {
		case lf.Rune:
			r2, err := readrune(runeScanner)
			switch {
			case io.EOF == err:
				// Nothing here.
			case cr.Rune == r2:
				_, err := runeWriter.WriteRune(r2)
				if nil != err {
					return erorr.Errorf("ini: problem writing rune: %w", err)
				}
			default:
				err := runeScanner.UnreadRune()
				if nil != err {
					return erorr.Errorf("ini: problem unreading rune: %w", err)
				}
			}
			return nil
		case cr.Rune:
			r2, err := readrune(runeScanner)
			switch {
			case io.EOF == err:
				// Nothing here.
			case lf.Rune == r2:
				_, err := runeWriter.WriteRune(r2)
				if nil != err {
					return erorr.Errorf("ini: problem writing rune: %w", err)
				}
			default:
				err := runeScanner.UnreadRune()
				if nil != err {
					return erorr.Errorf("ini: problem unreading rune: %w", err)
				}
			}
			return nil
		case nel.Rune, ls.Rune:
			return nil
		}
	}
}
*/
