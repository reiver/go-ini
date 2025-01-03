package ini

import (
	"io"
)

// AppendSequenceHeader appends the INI representation of a section header.
//
// You might use this function is you are implementing the [Marshaler] interface for a custom type.
// And in particular, use it (directly or indirectly) from you MarshalINI method.
//
// In INI, a sequence is used to create something similar to database records.
// For example:
//
//	[[fruits]]
//	
//	apple = ONE
//	Banana = TWO
//	CHERRY = THREE
//	dAtE = FOUR
//	
//	[[fruits]]
//	
//	apple = 1
//	Banana = 2
//	CHERRY = 3
//	dAtE = 4
//	
//	[[fruits]]
//	
//	apple = once
//	Banana = twice
//	CHERRY = thrice
//	dAtE = fource
//
// Also see [SequenceHeaderToString] and [WriteSequenceHeader]
func AppendSequenceHeader(p []byte, name ...string) []byte {
	if len(name) <= 0 {
		return p
	}

	p = append(p, "[["...)
	for index, part := range name {
		if 0 < index {
			p = append(p, '.')
		}
		p = EncodeAndAppendKey(p, part)
	}
	p = append(p, "]]"...)
	p = append(p, '\n')

	return p
}

// See also [AppendSequenceHeader] and [WriteSequenceHeader]
func SequenceHeaderToString(name ...string) string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = AppendSequenceHeader(p, name...)

	return string(p)
}

// See also [AppendSequenceHeader] and [SequenceHeaderToString]
func WriteSequenceHeader(dst io.Writer, name ...string) error {
	if nil == dst {
		return errNilWriter
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = AppendSequenceHeader(p, name...)

	_, err := dst.Write(p)
	if nil != err {
		return err
	}

	return nil
}
