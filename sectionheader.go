package ini

import (
	"io"
)

// AppendSectionHeader appends the INI representation of a section header.
//
// You might use this function is you are implementing the [Marshaler] interface for a custom type.
// And in particular, use it (directly or indirectly) from you MarshalINI method.
//
// In INI, a section used to create something to a single records. For example:
//
//	[basket]
//
//	apple = 1
//	Banana = 2
//	CHERRY = 3
//	dAtE = 4
//
// And in this example, the section-header is:
//
//	[basket]
//
// Which could be created with a call similar to:
//
//	p = ini.AppendSectionHeader(p, "basket")
//
// For another example, this call:
//
//	p = ini.AppendSectionHeader(p, "yek", "do", "se")
//
// Would producde the section-header:
//
//	[yek.do.se]
//
// Also see [SectionHeaderToString] and [WriteSectionHeader]
//
// Also, AppendSectionHeader shouldn't be confused with [AppendSequenceHeader]
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
