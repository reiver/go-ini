package ini

import (
	"strings"

	"github.com/reiver/go-erorr"
)

const (
	errINISetterCannotHandleSequenceHeader = erorr.Error("ini: etter cannot handle sequence-header")
)

type internalSetterPublisher struct {
	setter Setter
	section []string
}

var _ Publisher = &internalSetterPublisher{}

func (internalSetterPublisher) PublishINIComment(comment string) error {
	return nil
}

func (receiver *internalSetterPublisher) PublishINISectionHeader(section ...string) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.section = section
	return nil
}

func (receiver *internalSetterPublisher) PublishINISequenceHeader(...string) error {
	if nil == receiver {
		return errNilReceiver
	}

	return errINISetterCannotHandleSequenceHeader
}

func (receiver internalSetterPublisher) PublishINIKeyValue(name string, value string) error {
	var setter Setter = receiver.setter
	if nil == setter {
		return errNilSetter
	}

	var section []string = receiver.section

	var fullname string = name
	if 0 < len(section) {
		fullname = strings.Join(section, ".") + "." + fullname
	}

	return setter.SetININameValue(fullname, value)
}
