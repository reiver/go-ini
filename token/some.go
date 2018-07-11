package initoken

import (
	"io"
)

type SomeType interface {
	Type

	INITokenSome() SomeType

	String() string
	WriteTo(io.Writer) (int64, error)
}
