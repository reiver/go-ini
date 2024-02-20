package ini

type Publisher interface {
	PublishINIComment(comment string) error
	PublishINISection(section string) error
	PublishININameValue(name string, value string) error
}

