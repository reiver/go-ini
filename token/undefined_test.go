package initoken

import (
	"fmt"

	"testing"
)

func TestUndefinedString(t *testing.T) {

	var value string = "# This is a comment."

	var token Type = SomeUndefined(value)

	expected := fmt.Sprintf("initoken.SomeUndefined(%q)", value)

	if actual := token.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
}
