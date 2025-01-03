package ini

// Marshal returns the INI encoding of `value`.
//
// See also [ToString] and [Write]
func Marshal(value any) ([]byte, error) {
	return ContentOf(value)
}
