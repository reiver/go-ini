package ini

type internalMapKeyValueIterator struct {
	value map[string]string
}

func (receiver internalMapKeyValueIterator) For(fn func(string,string)error) error {
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
		var value string = receiver.value[key]
		err := fn(key, value)
		if nil != err {
			return err
		}
	}

	return nil
}
