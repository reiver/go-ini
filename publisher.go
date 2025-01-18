package ini

type Publisher interface {
	PublishINIComment(string) error
	PublishINIKeyValue(string, string) error
	PublishINISectionHeader(...string) error
	PublishINISequenceHeader(...string) error
}
