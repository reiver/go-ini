package ini

// Valuer is used by custom types to return their INI value.
//
// Valuer is used by [ValueOf]
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
type Valuer interface {
	INIValue() (string, error)
}
