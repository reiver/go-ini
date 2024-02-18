package iniscanner_section

// You can tell whether a line in an INI file is a section just based on the first character.
//
// If the first character is a '[', then it is a section.
//
// IsSection returns true if 'magic' is a INI file section, and returns false otherwise.
func IsSection(r rune) bool {
	switch r {
	case '[':
		return true
	default:
		return false
	}
}
