package ininamevalue

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errEmptyString    = erorr.Error("ini: empty string")
	errInternalError  = erorr.Error("ini: internal error")
	errNilRuneScanner = erorr.Error("ini: nil rune scanner")
	errNilRuneWriter  = erorr.Error("ini: nil rune writer")
	errRuneError      = erorr.Error("ini: rune error")
)
