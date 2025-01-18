package initest

type Publisher struct {
	PublishINICommentFunc        func(string) error
	PublishINIKeyValueFunc       func(string, string) error
	PublishINISectionHeaderFunc  func(...string) error
	PublishINISequenceHeaderFunc func(...string) error
}

func (receiver Publisher) PublishINIComment(comment string) error {
	return receiver.PublishINICommentFunc(comment)
}

func (receiver Publisher) PublishINIKeyValue(key string, value string) error {
	return receiver.PublishINIKeyValueFunc(key, value)
}

func (receiver Publisher) PublishINISectionHeader(name ...string) error {
	return receiver.PublishINISectionHeaderFunc(name...)
}

func (receiver Publisher) PublishINISequenceHeader(name ...string) error {
	return receiver.PublishINISequenceHeaderFunc(name...)
}
