package ini

type internalMapWrapper interface {
	AppendINIContent([]byte,...string) ([]byte, error)
}
