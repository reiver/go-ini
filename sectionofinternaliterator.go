package ini

func sectionOfInternalIterator(iterator internalKeyValueIterator) ([]byte, error) {

	if nil == iterator {
		return nil, errNilKeyValueIterator
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]
	{
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
