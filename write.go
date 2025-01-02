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
		var buffer [256]byte
		var p []byte = buffer[0:0]

		p = EncodeAndAppendKey(p, key)
		p = append(p, " = "...)
		p = EncodeAndAppendValue(p, value)
		p = append(p, '\n')

		_, err := dst.Write(p)
		if nil != err {
			return err
		}

		return nil
	})
	if nil != err {
		return err
	}

	return nil
}

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
