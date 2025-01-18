package ini

import (
	"bytes"
	"strings"

	"github.com/reiver/go-erorr"
)

const (
	errINISetterCannotHandleSequenceHeader = erorr.Error("ini: setter cannot handle sequence-header")
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

func (receiver *internalSetterPublisher) PublishINISequenceHeader(name ...string) error {
	if nil == receiver {
		return errNilReceiver
	}

	var errorMessage string
	{
		var buffer [256]byte
		var p []byte = buffer[0:0]

		p = append(p, "ini: setter cannot handle sequence-headers — there was attempt to set sequence header "...)
		p = AppendSequenceHeader(p, name...)
		p = bytes.TrimRight(p, "\n\r")

		errorMessage = string(p)
	}

	return erorr.Error(errorMessage)
}

func (receiver *internalSetterPublisher) PublishINIKeyValue(name string, value string) error {
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
