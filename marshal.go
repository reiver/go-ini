package ini

// Marshal returns the INI encoding of `content`.
//
// Some Go built-in types have an ini-content represention; such as map[string]any and map[string]string
//
// A custom type can also have an ini-content by implementing the [Marshaler] interface.
//
// See also [Append], [ToString] and [Write]
//
// Example of calling Marshal without a nesting:
//
//	bytes, err := ini.Marshal(data)
//
// Example of calling Marshal with a nesting:
//
//	bytes, err := ini.Marshal(data, "dairy", "milk")
func Marshal(content any, nesting ...string) ([]byte, error) {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	return Append(p, content, nesting...)
}
