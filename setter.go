package ini

type Setter interface {
	SetININameValue(name string, value string) error
}
