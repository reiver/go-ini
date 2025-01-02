package ini

// KeyValuer is used by custom types to return its INI value that can be represented by set of key-value pairs.
//
// See also [Valuer]
type KeyValuer interface {
	For(func(string,string)error)error
}
