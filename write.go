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

// Write writes the INI encoding of `src` to `dst`.
//
// See also [Marshal] and [ToString]
func Write(dst io.Writer, src any) error {
	if nil == dst {
		return errNilWriter
	}

	var keyvalueiter KeyValueIterator
	switch casted := src.(type) {
	case Marshaler:
		return WriteMarshaler(dst, casted)
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

// WriteMarshaler is similar to [Write] except that it specifically writes the INI for a [Marshaler] rather than any.
func WriteMarshaler(dst io.Writer, src Marshaler) error {
	if nil == dst {
		return errNilWriter
	}
	if nil == src {
		return errNilMarshaler
	}

	p, err := src.MarshalINI()
	if nil != err {
		return err
	}

	_, err = dst.Write(p)
	if nil != err {
		return err
	}

	return nil
}
