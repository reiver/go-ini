package initoken

import (
	"fmt"

	"testing"
)

func TestValueString(t *testing.T) {

	var value string = "value ۰, ۱, ۲, ۳, ۴, ۵, ۶, ۷, ۸, ۹"

	var token Type = SomeValue(value)

	expected := fmt.Sprintf("initoken.SomeValue(%q)", value)

	if actual := token.String(); expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
}
