package ini

// Marshaler is used by custom types to return its INI value.
//
// Behind the scenes, Marshaler is used by [ContentOf]
type Marshaler interface {
	MarshalINI([]byte, ...string) ([]byte, error)
}
