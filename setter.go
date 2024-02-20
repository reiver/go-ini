package ini

type Setter interface {
	SetINI(name string, value string) error
}

