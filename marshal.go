package ini

// Marshal returns the INI encoding of `value`.
//
// See also [ToString] and [Write]
//
// Notes that, Marshal is a wrapper around [ContentOf] with the 'nesting' parameter being empty.
func Marshal(value any) ([]byte, error) {
	return ContentOf(value)
}
