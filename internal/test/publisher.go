package initest

type Publisher struct {
	PublishINICommentFunc        func(string) error
	PublishINIKeyValueFunc       func(string, string) error
	PublishINISectionHeaderFunc  func(...string) error
	PublishINISequenceHeaderFunc func(...string) error
}

func (receiver Publisher) PublishINIComment(comment string) error {
	if nil == receiver.PublishINICommentFunc {
		return nil
	}
	return receiver.PublishINICommentFunc(comment)
}

func (receiver Publisher) PublishINIKeyValue(key string, value string) error {
	if nil == receiver.PublishINIKeyValueFunc {
		return nil
	}
	return receiver.PublishINIKeyValueFunc(key, value)
}

func (receiver Publisher) PublishINISectionHeader(name ...string) error {
	if nil == receiver.PublishINISectionHeaderFunc {
		return nil
	}
	return receiver.PublishINISectionHeaderFunc(name...)
}

func (receiver Publisher) PublishINISequenceHeader(name ...string) error {
	if nil == receiver.PublishINISequenceHeaderFunc {
		return nil
	}
	return receiver.PublishINISequenceHeaderFunc(name...)
}
