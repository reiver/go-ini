package ini

type internalMapStringSliceString struct {
	value map[string][]string
}

func (receiver internalMapStringSliceString) IsEmpty() bool {
	return  len(receiver.value) <= 0
}

func (receiver internalMapStringSliceString) MarshalINI(p []byte, nesting ...string) ([]byte, error) {
	if receiver.IsEmpty() {
		return p, nil
	}

	if 0 < len(nesting) {
		p = AppendSectionHeader(p, nesting...)
	}

	err := receiver.ForEachKeyStringValue(func(key string, value string) error {
		p = AppendKeyValue(p, key, value)
		return nil
	})
	if nil != err {
		return p, err
	}

	return p, nil
}

func (receiver internalMapStringSliceString) Keys() []string {
	return mapKeys(receiver.value)
}

func (receiver internalMapStringSliceString) ForEachKeyMapValue(fn func(string,Marshaler)error) error {
	return nil
}

func (receiver internalMapStringSliceString) ForEachKeyStringValue(fn func(string,string)error) error {
	if receiver.IsEmpty() {
		return nil
	}

	var keys []string = receiver.Keys()

	for _, key := range keys {
		var values []string = receiver.value[key]

		for _, value := range values {
			err := fn(key, value)
			if nil != err {
				return err
			}
		}
	}

	return nil
}
