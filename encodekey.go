package ini

// EncodeAndAppendKey is similar to [EncodeKey] except it appends the result to a []byte.
func EncodeAndAppendKey(p []byte, key string) []byte {
	var length int = len(key)

	for index:=0; index<length; index++ {
		var b byte = key[index]

		switch b {
		case '"':  // U+0022
			p = append(p, `\"`...)
		case '#':  // U+0022
			p = append(p, `\#`...)
		case '\'': // U+0027
			p = append(p, `\'`...)
		case ':':  // U+003A
			p = append(p, `\:`...)
		case '=':  // U+003D
			p = append(p, `\=`...)
		case ';':  // U+003B
			p = append(p, `\;`...)
		case '\\': // U+009C
			p = append(p, `\\`...)
		default:
			p = append(p, b)
		}
	}

	return p
}

// EncodeKey (potentially) encodes a key (from a key-value pair) to make it valid within INI content.
func EncodeKey(key string) string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = EncodeAndAppendKey(p, key)

	return string(p)
}
