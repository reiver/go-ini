package initoken

import (
	"fmt"

	"testing"
)

func TestSectionEndString(t *testing.T) {

	var value string = "}"

	var token Type = SomeSectionEnd(value)

	expected := fmt.Sprintf("initoken.SomeSectionEnd(%q)", value)

	if actual := token.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
}
