package ini

import (
	"github.com/reiver/go-erorr"
)

// Append is similar to [Marshal] except that it appends.
//
// See also [Marshal], [ToString] and [Write]
func Append(p []byte, content any, nesting ...string) ([]byte, error) {
	var marshaler Marshaler

	switch casted := content.(type) {
	case Marshaler:
		marshaler = casted
	case map[string]string:
		marshaler = internalMapStringString{casted}
	case map[string][]string:
		marshaler = internalMapStringSliceString{casted}
	case map[string]any:
		marshaler = internalMapStringAny{casted}
	case []map[string]string:
		if len(nesting) <= 0 {
			return nil, erorr.Errorf("ini: type %T must be nested", content)
		}
		marshaler = internalSliceMapStringString{casted}
	default:
		return nil, erorr.Errorf("ini: type %T does not have a 'content' representation", content)
	}

	return marshaler.MarshalINI(p, nesting...)
}
