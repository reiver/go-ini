package ini

type internalSectionerAsContenter struct {
	sectioner Sectioner
}

var _ Contenter = internalSectionerAsContenter{}

func (receiver internalSectionerAsContenter) AppendINIContent(p []byte, nesting ...string) ([]byte, error) {
	return receiver.sectioner.AppendINISection(p, nesting...)
}
