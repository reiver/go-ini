package ini

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilBytes      = erorr.Error("ini: nil bytes")
	errInternalError = erorr.Error("ini: internal error")
	errNilMarshaler  = erorr.Error("ini: nil marshaler")
	errNilMap        = erorr.Error("ini: nil map")
	errNilPublisher  = erorr.Error("ini: nil publisher")
	errNilReceiver   = erorr.Error("ini: nil receiver")
	errNilSetter     = erorr.Error("ini: nil setter")
	errNilWriter     = erorr.Error("ini: nil writer")
	errRuneError     = erorr.Error("ini: rune error")
)
