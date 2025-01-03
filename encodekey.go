package ini

import (
	"github.com/reiver/go-ascii"
	"github.com/reiver/go-unicode"
)

// encodeAndAppendKey is similar to [encodeKey] except it appends the result to a []byte.
func encodeAndAppendKey(p []byte, key string) []byte {
	const NELstr string = string(unicode.NEL)
	const LSstr  string = string(unicode.LS)
	const PSstr  string = string(unicode.PS)

	for _, r := range key {

		switch r {
		case ascii.LF: // U+000A
			p = append(p, '\\')
			p = append(p, ascii.LF)
		case ascii.VT: // U+000B
			p = append(p, '\\')
			p = append(p, ascii.VT)
		case ascii.FF: // U+000C
			p = append(p, '\\')
			p = append(p, ascii.FF)
		case ascii.CR: // U+000D
			p = append(p, '\\')
			p = append(p, ascii.CR)
		case '"':  // U+0022
			p = append(p, `\"`...)
		case '#':  // U+0022
			p = append(p, `\#`...)
		case '\'': // U+0027
			p = append(p, `\'`...)
		case '.': // U+002E
			p = append(p, `\.`...)
		case ':':  // U+003A
			p = append(p, `\:`...)
		case '=':  // U+003D
			p = append(p, `\=`...)
		case ';':  // U+003B
			p = append(p, `\;`...)
		case '[':  // U+005B
			p = append(p, `\[`...)
		case ']':  // U+005D
			p = append(p, `\]`...)
		case '\\': // U+009C
			p = append(p, `\\`...)
		case unicode.NEL: // U+0085
			p = append(p, '\\')
			p = append(p, NELstr...)
		case unicode.LS: // U+2028
			p = append(p, '\\')
			p = append(p, LSstr...)
		case unicode.PS: // U+2029
			p = append(p, '\\')
			p = append(p, PSstr...)
		default:
			p = append(p, string(r)...)
		}
	}

	return p
}

// encodeKey (potentially) encodes a key (from a key-value pair) to make it valid within INI content.
func encodeKey(key string) string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = encodeAndAppendKey(p, key)

	return string(p)
}
