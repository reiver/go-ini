package ini

import (
	"io"
)

// Encoder is the interface implemented by types that can encode themselves into valid INI.
type Encoder interface {
	EncodeINI(io.Writer) error
}
