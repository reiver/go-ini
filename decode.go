package ini

import (
	"io"
	"strings"

	"github.com/reiver/go-runewriter"
	"sourcecode.social/reiver/go-erorr"

	"github.com/reiver/go-ini/internal/comment"
	"github.com/reiver/go-ini/internal/namevalue"
	"github.com/reiver/go-ini/internal/section"
)

func Decode(dst any, runescanner io.RuneScanner) error {

	if nil == runescanner {
		return errNilRuneScanner
	}

	var result *map[string][]string
	{
		var casted bool

		result, casted = dst.(*map[string][]string)
		if !casted {
			return erorr.Errorf("ini: cannot decode into %T", dst)
		}
	}

	var section string = ""

	r, _, err := runescanner.ReadRune()
	if nil != err {
		return erorr.Errorf("ini: problem reading rune: %w", err)
	}

	switch {
	case inicomment.IsMagic(r):
		err := runescanner.UnreadRune()
		if nil != err {
			return erorr.Errorf("ini: problem un-reading rune: %w", err)
		}

		err = inicomment.Copy(runewriter.Discard, runescanner)
		if nil != err {
			return erorr.Errorf("ini: problem reading ini comment: %w", err)
		}
	case inisection.IsMagic(r):
		err := runescanner.UnreadRune()
		if nil != err {
			return erorr.Errorf("ini: problem un-reading rune: %w", err)
		}

		var runeWriter strings.Builder

		err = inisection.Copy(&runeWriter, runescanner)
		if nil != err {
			return erorr.Errorf("ini: problem reading ini section: %w", err)
		}

		section, err = inisection.ParseString(runeWriter.String())
		if nil != err {
			return err
		}
	default:
		var runeWriter strings.Builder

		err := ininamevalue.Copy(&runeWriter, runescanner)
		if nil != err {
			return erorr.Errorf("ini: problem reading ini name-value pair: %w", err)
		}

		name, value, err := ininamevalue.ParseString(runeWriter.String())

		var fullname string = name
		if "" != section {
			fullname = section + "." + name
		}

		(*result)[fullname] = []string{value}
	}

	return nil
}
