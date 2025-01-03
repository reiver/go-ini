package ini

type internalKeyValueIterator interface {
	For(fn func(string,string)error) error
	Sub(fn func(internalKeyValueIterator)error) error
}

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
	{
		for key, _ := range receiver.value {
			keys = append(keys, key)
		}

		SortKeys(keys)
	}

	for _, key := range keys {
		var value T = receiver.value[key]

		var stringvalue string
		{
			var err error

			stringvalue, err = ValueOf(value)
			if nil != err {
	/////////////////////// CONTINUE
				continue
			}
		}

		err := fn(key, stringvalue)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver internalMapKeyValueIterator[T]) Sub(fn func(internalKeyValueIterator)error) error {
	if nil == fn {
		return nil
	}
	if len(receiver.value) <= 0 {
		return nil
	}

	var keys []string
	{
		for key, _ := range receiver.value {
			keys = append(keys, key)
		}

		SortKeys(keys)
	}

	for _, key := range keys {
		var value T = receiver.value[key]

		var something any = value

		var iterator internalKeyValueIterator
		{
			switch casted := something.(type) {
			case map[string]string:
				iterator = internalMapKeyValueIterator[string]{casted}
			case map[string]any:
				iterator = internalMapKeyValueIterator[any]{casted}
			default:
	/////////////////////// CONTINUE
				continue
			}
		}

		err := fn(iterator)
		if nil != err {
			return err
		}
	}

	return nil
}
