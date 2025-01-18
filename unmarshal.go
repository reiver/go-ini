package ini

import (
	"unicode/utf8"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-ini/internal/comment"
	"github.com/reiver/go-ini/internal/namevalue"
	"github.com/reiver/go-ini/internal/section"
	"github.com/reiver/go-ini/internal/sequence"
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
			var ahead []byte = p[inisection.MagicSize:]
			if len(ahead) <= 0 {
				return erorr.Error("ini: both INI 'section' and 'sequence' must have a 2nd character")
			}

			peek, size := utf8.DecodeRune(ahead)
			if utf8.RuneError == peek {
				switch size {
				case 1:
					return errRuneError
				default:
					return errInternalError
				}
			}

			switch {
			case inisequence.Is2ndMagic(peek):
				name, size, err := inisequence.ParseBytes(p)
				if nil != err {
					return erorr.Errorf("ini: problem reading ini sequence: %w", err)
				}

				if err := publisher.PublishINISequenceHeader(name...); nil != err {
					return erorr.Errorf("ini: problem publishing INI sequence: %w", err)
				}
				p = p[size:]
			default:
				name, size, err := inisection.ParseBytes(p)
				if nil != err {
					return erorr.Errorf("ini: problem reading ini section: %w", err)
				}

				if err := publisher.PublishINISectionHeader(name...); nil != err {
					return erorr.Errorf("ini: problem publishing INI section: %w", err)
				}
				p = p[size:]
			}
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
