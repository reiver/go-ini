package ini

// Sectioner is used by custom types to return its INI value that can be represented by as a single INI section.
//
// Behind the scenes, Sectionr is used by [SectionOf]
//
// For example:
//
//	type MyStruct struct {
//		// ..
//	}
//	
//	func (recevier MyStruct) INISection() (string, error) {
//		// ...
//	}
//
// For custom types that would be represented by more than just a simple INI value, see [KeySectionr]
type Sectioner interface {
	INISection() ([]byte, error)
}
