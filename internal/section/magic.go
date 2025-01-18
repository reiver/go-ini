package inisection

const magic rune = '['

// measured in bytes
const MagicSize = len(string(magic))

// You can tell whether a line in an INI file is a section just based on the first character.
//
// If the first character is a '[', then it is a section.
//
// IsMagic returns true if 'magic' is a INI file section, and returns false otherwise.
func IsMagic(r rune) bool {
	switch r {
	case magic:
		return true
	default:
		return false
	}
}
