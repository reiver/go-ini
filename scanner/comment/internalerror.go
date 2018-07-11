package iniscanner_comment

import (
	"fmt"
)

type InternalErrorComplainer interface {
	error
	Err()
	InternalErrorComplainer()
}

type internalInternalErrorComplainer struct{
	value    string
	activity string
	err      error
}

func (complainer internalInternalErrorComplainer) Error() string {
	return fmt.Sprintf("Internal error parsing comment %q when: %s; because: %s", complainer.value, complainer.activity, complainer.err.Error())
}

func (complainer internalInternalErrorComplainer) Err() error {
	return complainer.err
}

func (complainer internalInternalErrorComplainer) InternalErrorComplainer() {
	// Nothing here.
}
