package ini

type internalMapStringString struct {
	value map[string]string
}

func (receiver internalMapStringString) IsEmpty() bool {
	return  len(receiver.value) <= 0
}

func (receiver internalMapStringString) MarshalINI(p []byte, nesting ...string) ([]byte, error) {
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

func (receiver internalMapStringString) Keys() []string {
	return mapKeys(receiver.value)
}

func (receiver internalMapStringString) ForEachKeyMapValue(fn func(string,Marshaler)error) error {
	return nil
}

func (receiver internalMapStringString) ForEachKeyStringValue(fn func(string,string)error) error {
	if receiver.IsEmpty() {
		return nil
	}

	var keys []string = receiver.Keys()

	for _, key := range keys {
		var value string = receiver.value[key]

		err := fn(key, value)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver internalMapStringString) INIContent(nesting ...string) ([]byte, error) {
	if receiver.IsEmpty() {
		return nil, nil
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	return receiver.MarshalINI(p, nesting...)
}
