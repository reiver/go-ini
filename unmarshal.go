package ini

import (
	"unicode/utf8"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-ini/internal/comment"
	"github.com/reiver/go-ini/internal/namevalue"
	"github.com/reiver/go-ini/internal/section"
)

// Unmarshal parses the INI-encoded data and stores the result in the value pointed to by 'dst'.
// If 'dst' is nil or not a pointer, Unmarshal returns an [InvalidDestinationTypeError].
func Unmarshal(data []byte, dst any) error {

	if nil == data {
		return errNilBytes
	}
	if nil == dst {
		return errNilDestination
	}

	var publisher Publisher
	{
		switch casted := dst.(type) {
		case Publisher:
			publisher = casted
		case Setter:
			publisher = &internalSetterPublisher{
				setter:casted,
			}
		case *map[string]string:
			if nil == *casted {
				m := make(map[string]string)
				*casted = m
			}
			setter := internalMapStringStringSetter{
				mapStringString: casted,
			}
			publisher = &internalSetterPublisher{
				setter:setter,
			}
//		case *map[string][]string:
//			publisher = 
//		case *map[string]any:
//			publisher = &internalMapStringAnyPublisher{casted}
		default:
			return CreateInvalidDestinationTypeError(dst)
		}
	}

	return unmarshal(data, publisher)
}

func unmarshal(data []byte, publisher Publisher) error {

	if nil == data {
		return errNilBytes
	}
	if nil == publisher {
		return errNilPublisher
	}

	var p []byte = data

	for {
		if len(p) <= 0 {
			break
		}

		var r rune
		{
			var size int

			r, size = utf8.DecodeRune(p)
			if utf8.RuneError == r {
				switch size {
				case 1:
                        	        return errRuneError
				default:
        	                        return errInternalError
				}
			}
		}

		switch {
		case inicomment.IsMagic(r):
			comment, size, err := inicomment.ParseBytes(p)
			if nil != err {
				return erorr.Errorf("ini: problem reading ini comment: %w", err)
			}

			if err := publisher.PublishINIComment(comment); nil != err {
				return erorr.Errorf("ini: problem publishing INI comment: %w", err)
			}
			p = p[size:]
		case inisection.IsMagic(r):
			section, size, err := inisection.ParseBytes(p)
			if nil != err {
				return erorr.Errorf("ini: problem reading ini section: %w", err)
			}

			if err := publisher.PublishINISectionHeader(section); nil != err {
				return erorr.Errorf("ini: problem publishing INI section: %w", err)
			}
			p = p[size:]
		default:
			name, value, size, err := ininamevalue.ParseBytes(p)
			if nil != err {
				return erorr.Errorf("ini: problem reading ini key-value: %w", err)
			}

			if "" != name {
				if err := publisher.PublishINIKeyValue(name, value); nil != err {
					return erorr.Errorf("ini: problem publishing INI key-value: %w", err)
				}
			}
			p = p[size:]
		}
	}

	return nil
}
