package ini

// Contenter is used by custom types to return its INI value.
//
// Behind the scenes, Contenter is used by [ContentOf]
type Contenter interface {
	AppendINIContent([]byte, ...string) ([]byte, error)
}
