package ini

type internalMapStringAny struct {
	value map[string]any
}

func (receiver internalMapStringAny) IsEmpty() bool {
	return  len(receiver.value) <= 0
}

func (receiver internalMapStringAny) MarshalINI(p []byte, nesting ...string) ([]byte, error) {

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

	err = receiver.ForEachKeyMapValue(func(key string, marshaler Marshaler) error {
		if nil == marshaler {
			return errNilMarshaler
		}

		var newNesting []string = append([]string(nil), nesting...)
		newNesting = append(newNesting, key)

		var e error
		p, e = marshaler.MarshalINI(p, newNesting...)
		if nil != e {
			return e
		}

		return nil
	})
	if nil != err {
		return p, err
	}

	return p, nil
}

func (receiver internalMapStringAny) Keys() []string {
	return mapKeys(receiver.value)
}

func (receiver internalMapStringAny) ForEachKeyMapValue(fn func(string,Marshaler)error) error {
	if receiver.IsEmpty() {
		return nil
	}

	var keys []string = receiver.Keys()

	for _, key := range keys {
		var valueAny any = receiver.value[key]

		var value Marshaler

		switch casted := valueAny.(type) {
		case map[string]string:
			value = internalMapStringString{casted}
		case map[string]any:
			value = internalMapStringAny{casted}
		default:
	/////////////// CONTINUE
			continue
		}

		err := fn(key, value)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver internalMapStringAny) ForEachKeyStringValue(fn func(string,string)error) error {
	if receiver.IsEmpty() {
		return nil
	}

	var keys []string = receiver.Keys()

	for _, key := range keys {
		var valueAny any = receiver.value[key]

		var valueString string
		var err error

		valueString, err = ValueOf(valueAny)
		if nil != err {
	/////////////// CONTINUE
			continue
		}

		err = fn(key, valueString)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver internalMapStringAny) INIContent(nesting ...string) ([]byte, error) {

	if receiver.IsEmpty() {
		return nil, nil
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	return receiver.MarshalINI(p, nesting...)
}
