package ini

import (
	"github.com/reiver/go-humane"
)

// SortKeys sorts a record's keys in a human-friendly and humane way.
func SortKeys(keys []string) {
	humane.SortStrings(keys)
}
