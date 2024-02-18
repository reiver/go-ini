package iniscanner_comment

func Peek(r rune) bool {
	switch r {
	case ';','#':
		return true
	default:
		return false
	}
}
