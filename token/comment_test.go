package initoken

import (
	"fmt"

	"testing"
)

func TestCommentString(t *testing.T) {

	var value string = "# This is a comment."

	var token Type = SomeComment(value)

	expected := fmt.Sprintf("initoken.SomeComment(%q)", value)

	if actual := token.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
}
