package iniscanner_value

import (
	"github.com/reiver/go-whitespace"
)

func Peek(r rune) bool {
	return ! whitespace.IsWhitespace(r)
}
