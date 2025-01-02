package ini

import (
	"io"

	"github.com/reiver/go-erorr"
)

const (
	errNilKeyValueIterator = erorr.Error("ini: nil key-valuer")
)

func write(dst io.Writer, src KeyValueIterator) error {

	if nil == dst {
		return errNilWriter
	}

	err := src.For(func(key string, value string) error {
		return WriteKeyValue(dst, key, value)
	})
	if nil != err {
		return err
	}

	return nil
}

// See also [Marshal]
func Write(dst io.Writer, src any) error {

	if nil == dst {
		return errNilWriter
	}

	var keyvalueiter KeyValueIterator
	switch casted := src.(type) {
	case KeyValueIterator:
	keyvalueiter = casted
	case map[string]string:
		keyvalueiter = internalMapKeyValueIterator[string]{casted}
	default:
		return erorr.Errorf("ini: cannot write-ini for something of type %T", src)
	}
	if nil == keyvalueiter {
		return errNilKeyValueIterator
	}

	err := write(dst, keyvalueiter)
	if nil != err {
		return err
	}

	return nil
}
