package ini

type Marshaler interface {
	MarshalINI(...string) ([]byte, error)
}
