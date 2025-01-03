package ini

func contentOfInternalIterator(iterator internalKeyValueIterator, nesting ...string) ([]byte, error) {

	if nil == iterator {
		return nil, errNilKeyValueIterator
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	{
		var appended bool
		err := iterator.For(func(key string, value string) error {
			p = AppendKeyValue(p, key, value)
			appended = true
			return nil
		})
		if nil != err {
			return nil, err
		}
		if appended {
			p = append(p, '\n')
		}

		err = iterator.Sub(func(iterator internalKeyValueIterator, nesting ...string) error {
			p = AppendSectionHeader(p, nesting...)
			p = append(p, '\n')
			return nil
		})
		if nil != err {
			return nil, err
		}
	}

	return p, nil
}
