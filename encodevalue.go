package ini

// EncodeAndAppendValue is similar to [EncodeValue] except it appends the result to a []byte.
func EncodeAndAppendValue(p []byte, key string) []byte {
	return EncodeAndAppendKey(p, key)
}

// EncodeValue (potentially) encodes a value (from a key-value pair) to make it valid within INI content.
func EncodeValue(key string) string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = EncodeAndAppendValue(p, key)

	return string(p)
}
