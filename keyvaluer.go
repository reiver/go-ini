package ini

// See also [KeyValuer]
type KeyValuer interface {
	INIValue() (string, error)
}
