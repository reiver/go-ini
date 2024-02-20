package ini

type internalSetterPublisher struct {
	setter Setter
	section string
}

var _ Publisher = &internalSetterPublisher{}

func (internalSetterPublisher) PublishINIComment(comment string) error {
	return nil
}

func (receiver *internalSetterPublisher) PublishINISection(section string) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.section = section
	return nil
}

func (receiver internalSetterPublisher) PublishININameValue(name string, value string) error {
	var setter Setter = receiver.setter
	if nil == setter {
		return errNilSetter
	}

	var section string = receiver.section

	var fullname string = name
	if "" != section {
		fullname = section + "." + name
	}

	return setter.SetINI(fullname, value)
}
