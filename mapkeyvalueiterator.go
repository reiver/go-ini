package ini

import (
	"github.com/reiver/go-erorr"
)

type internalMapKeyValueIterator[T any] struct {
	value map[string]T
}

func (receiver internalMapKeyValueIterator[T]) For(fn func(string,string)error) error {
	if nil == fn {
		return nil
	}
	if len(receiver.value) <= 0 {
		return nil
	}

	var keys []string
	for key, _ := range receiver.value {
		keys = append(keys, key)
	}

	SortKeys(keys)

	for _, key := range keys {
		var value T = receiver.value[key]

		var stringvalue string
		{
			var err error

			stringvalue, err = ValueOf(value)
			if nil != err {
				return erorr.Errorf("ini: problem casting from 'string' to '%T': %w", value, err)
			}
		}

		err := fn(key, stringvalue)
		if nil != err {
			return err
		}
	}

	return nil
}
