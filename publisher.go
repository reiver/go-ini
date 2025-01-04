package ini

type Publisher interface {
	PublishINIComment(comment string) error
	PublishINISectionHeader(...string) error
	PublishINISequenceHeader(...string) error
	PublishININameValue(name string, value string) error
}
