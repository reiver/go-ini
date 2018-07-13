package initoken

import (
	"testing"
)

func TestNoneString(t *testing.T) {

	var token Type = None()

	expected := "initoken.None()"

	if actual := token.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
}
