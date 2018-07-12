package iniscanner_separator

func Peek(r rune) bool {
	switch r {
	case '=',':':
		return true
	default:
		return false
	}
}
