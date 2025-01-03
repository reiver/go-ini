package ini

import (
	"github.com/reiver/go-erorr"
)

const (
	errSequenceCannotHaveEmptyNesting = erorr.Error("ini: sequence cannot have empty nesting")
)

type internalSliceMapStringString struct {
	value []map[string]string
}

func (receiver internalSliceMapStringString) IsEmpty() bool {
	return  len(receiver.value) <= 0
}

func (receiver internalSliceMapStringString) MarshalINI(p []byte, nesting ...string) ([]byte, error) {
	if receiver.IsEmpty() {
		return p, nil
	}

	if len(nesting) <= 0 {
		return nil, errSequenceCannotHaveEmptyNesting
	}

	var headerBuffer [256]byte
	var header []byte = headerBuffer[0:0]
	header = AppendSequenceHeader(header, nesting...)

	for _, datum := range receiver.value {

		p = append(p, header...)

		var err error
		p, err = Append(p, datum)
		if nil != err {
			return p, err
		}
	}

	return p, nil
}
