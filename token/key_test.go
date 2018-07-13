package initoken

import (
	"fmt"

	"testing"
)

func TestKeyString(t *testing.T) {

	var value string = "key"

	var token Type = SomeKey(value)

	expected := fmt.Sprintf("initoken.SomeKey(%q)", value)

	if actual := token.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
}
