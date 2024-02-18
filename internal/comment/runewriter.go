package inicomment

// Why doesn't the Go built-in package "io" have this?!?!?
// It has io.RuneReader.
type RuneWriter interface {
	WriteRune(r rune) (size int, err error)
}
