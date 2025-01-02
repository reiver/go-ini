package ini

func sectionOfMap[T any](m map[string]T) ([]byte, error) {

	if len(m) <= 0 {
		return nil, nil
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]
	{
		iterator := internalMapKeyValueIterator[T]{m}

		err := iterator.For(func(key string, value string) error {
			p = AppendKeyValue(p, key, value)
			return nil
		})
		if nil != err {
			return nil, err
		}
	}


	return p, nil
}
