package initoken

import (
	"fmt"

	"testing"
)

func TestSeparatorString(t *testing.T) {

	var value string = "# This is a comment."

	var token Type = SomeSeparator(value)

	expected := fmt.Sprintf("initoken.SomeSeparator(%q)", value)

	if actual := token.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
}
