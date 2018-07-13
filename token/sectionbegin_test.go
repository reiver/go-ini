package initoken

import (
	"fmt"

	"testing"
)

func TestSectionBeginString(t *testing.T) {

	var value string = "{"

	var token Type = SomeSectionBegin(value)

	expected := fmt.Sprintf("initoken.SomeSectionBegin(%q)", value)

	if actual := token.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
}
