package inicomment

// You can tell whether a line in an INI file is a comment just based on the first character.
//
// If the first character is a ';' or a '#', then it is a comment.
//
// IsCommentMagic returns true if 'magic' is a INI file comment, and returns false otherwise.
func IsCommentMagic(magic rune) bool {
	switch magic {
	case '#',';':
		return true
	default:
		return false
	}
}
