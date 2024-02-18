package inisection

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errEmptyString    = erorr.Error("ini: empty string")
	errMagicNotFound  = erorr.Error("ini: section magic character (i.e., '[' (U+005B)) not found")
	errNilRuneScanner = erorr.Error("ini: nil rune scanner")
	errNilRuneWriter  = erorr.Error("ini: nil rune writer")
)
