package ini

// See also [Marshal] and [Write]
func ToString(value any) (string, error) {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	var err error

	p, err = Marshal(value)
	if nil != err {
		return "", err
	}

	return string(p), nil
}

// NestedToString is similar to [ToString] except nested.
//
// See also [NestedMarshal] and [NestedWrite]
func NestedToString(value any, nesting ...string) (string, error) {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	var err error

	p, err = NestedMarshal(value, nesting...)
	if nil != err {
		return "", err
	}

	return string(p), nil
}
