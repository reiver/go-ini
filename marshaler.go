package ini

// Marshaler is the interface implemented by types that can marshal themselves into valid INI.
type Marshaler interface {
	MarshalINI() ([]byte, error)
}
