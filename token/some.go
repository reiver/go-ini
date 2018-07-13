package initoken

import (
	"io"
)

type SomeType interface {
	Type

	INITokenSome() SomeType

	WriteTo(io.Writer) (int64, error)
}
