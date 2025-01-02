package ini

import (
	"io"

	"github.com/reiver/go-erorr"
)

const (
	errNilKeyValueIterator = erorr.Error("ini: nil key-valuer")
)

// NestedWrite is similar to [Write] except nested.
//
// See also [NestedMarshal] and [NestedToString]
func NestedWrite(dst io.Writer, src any, nesting ...string) error {
	if nil == dst {
		return errNilWriter
	}

	if 0 < len(nesting) {
		err := WriteSectionHeader(dst, nesting...)
		if nil != err {
			return err
		}
	}

	var keyvalueiter KeyValueIterator
	switch casted := src.(type) {
	case Marshaler:
		return writeMarshaler(dst, casted)
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
	return NestedWrite(dst, src)
}

func writeMarshaler(dst io.Writer, src Marshaler) error {
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
