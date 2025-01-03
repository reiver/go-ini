package ini

// Valuer is used by custom types to return its INI value that can be represented by as a simple INI value.
//
// Behind the scenes, Valuer is used by [ValueOf]
//
// For example:
//
//	type MyStruct struct {
//		// ..
//	}
//	
//	func (recevier MyStruct) INIValue() (string, error) {
//		// ...
//	}
//
// For custom types that would be represented by more than just a simple INI value, see [Marshaler]
type Valuer interface {
	INIValue() (string, error)
}
