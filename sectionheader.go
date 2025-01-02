package ini

import (
	"io"
)

// See also [SectionHeaderToString] and [WriteSectionHeader]
func AppendSectionHeader(p []byte, name ...string) []byte {
	if len(name) <= 0 {
		return p
	}

	p = append(p, '[')
	for index, part := range name {
		if 0 < index {
			p = append(p, '.')
		}
		p = EncodeAndAppendKey(p, part)
	}
	p = append(p, ']')
	p = append(p, '\n')

	return p
}

// See also [AppendSectionHeader] and [WriteSectionHeader]
func SectionHeaderToString(name ...string) string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = AppendSectionHeader(p, name...)

	return string(p)
}

// See also [AppendSectionHeader] and [SectionHeaderToString]
func WriteSectionHeader(dst io.Writer, name ...string) error {
	if nil == dst {
		return errNilWriter
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = AppendSectionHeader(p, name...)

	_, err := dst.Write(p)
	if nil != err {
		return err
	}

	return nil
}
