package ini

type internalMapStringStringSetter struct {
	mapStringString *map[string]string
}

var _ Setter = internalMapStringStringSetter{}

func (receiver internalMapStringStringSetter) SetININameValue(name string, value string) error {
	mapStringString := receiver.mapStringString
	if nil == mapStringString {
		return errNilMap
	}

	(*mapStringString)[name] = value
	return nil
}
