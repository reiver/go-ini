package initoken

import (
	"fmt"

	"testing"
)

func TestWhitespaceString(t *testing.T) {

	var value string = "  \r\n\r\n    \v   \u2028"

	var token Type = SomeWhitespace(value)

	expected := fmt.Sprintf("initoken.SomeWhitespace(%q)", value)

	if actual := token.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
}
