package ini

// DecodeKey (potentially) decodes a key (from a key-value pair) from its valid INI content form.
func DecodeKey(key string) string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	var length int = len(key)

	var escaped bool = false

	for index:=0; index<length; index++ {
		var b byte = key[index]

		switch {
		case escaped:
			p = append(p, b)
			escaped = false
		case '\\' == b:
			escaped = true
		default:
			p = append(p, b)
		}
	}

	return string(p)
}
