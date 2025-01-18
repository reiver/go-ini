package ini_test

import (
	"testing"

	"github.com/reiver/go-ini"
)

func TestUnmarshal_errNilDestination(t *testing.T) {

	var bytes []byte = []byte(`
[stuff]

apple  = once
Banana = twice
CHERRY = thrice
dAte   = fource
`)

	err := ini.Unmarshal(bytes, nil)

	if nil == err {
		t.Errorf("Expected an error but did not actually got one.")
		return
	}

	expected := "ini: nil destination"
	actual   := err.Error()

	if expected != actual {
		t.Errorf("The actual error-message is not what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", actual)
		return
	}
}
