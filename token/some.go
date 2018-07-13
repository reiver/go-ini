package initoken

import (
	"io"
)

type SomeType interface {
	Type

	INITokenSome() SomeType

	Unwrap() string
	WriteTo(io.Writer) (int64, error)
}
