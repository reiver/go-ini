package inisequence

import (
	"github.com/reiver/go-erorr"
)

const (
	errInternalError  = erorr.Error("ini: internal error")
	errNilBytes       = erorr.Error("ini: nil bytes")
	errRuneError      = erorr.Error("ini: rune error")
)
