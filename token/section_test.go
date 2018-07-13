package initoken

import (
	"fmt"

	"testing"
)

func TestSectionString(t *testing.T) {

	var value string = "[the section]"

	var token Type = SomeSection(value)

	expected := fmt.Sprintf("initoken.SomeSection(%q)", value)

	if actual := token.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
}
