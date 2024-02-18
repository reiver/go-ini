package inisection

import (
	"io"

	"sourcecode.social/reiver/go-erorr"
)

func readrune(runeScanner io.RuneScanner) (rune, error) {

	r, size, err := runeScanner.ReadRune()
	if io.EOF == err {
		return r, err
	}
	if nil != err {
		return r, erorr.Errorf("ini: problem reading rune in section: %w", err)
	}
	if size <= 0 {
		return r, erorr.Error("ini: problem reading rune in section")
	}

	return r, err
}
