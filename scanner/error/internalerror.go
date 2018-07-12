package iniscanner_error

import (
	"fmt"
)

type InternalErrorComplainer interface {
	error
	Err()
	InternalErrorComplainer()
}

func InternalError(value string, activity string, err error) error {
	return internalInternalErrorComplainer{
		err: err,
		activity: activity,
		value: value,
	}
}

type internalInternalErrorComplainer struct{
	value    string
	activity string
	err      error
}

func (complainer internalInternalErrorComplainer) Error() string {
	return fmt.Sprintf("Internal error parsing %q when: %s; because: %s", complainer.value, complainer.activity, complainer.err.Error())
}

func (complainer internalInternalErrorComplainer) Err() error {
	return complainer.err
}

func (complainer internalInternalErrorComplainer) InternalErrorComplainer() {
	// Nothing here.
}
