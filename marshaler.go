package ini

type Marshaler interface {
	MarshalINI() ([]byte, error)
}
