package ininamevalue

import (
	"io"

	"github.com/reiver/go-runewriter"

	"sourcecode.social/reiver/go-eol/cr"
	"sourcecode.social/reiver/go-eol/lf"
	"sourcecode.social/reiver/go-eol/ls"
	"sourcecode.social/reiver/go-eol/nel"
	"sourcecode.social/reiver/go-erorr"
)

func Copy(runeWriter runewriter.RuneWriter, runeScanner io.RuneScanner) error {
	if nil == runeScanner  {
		return errNilRuneScanner
	}
	if nil == runeWriter {
		return errNilRuneWriter
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