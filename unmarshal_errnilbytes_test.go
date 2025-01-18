package ini_test

import (
	"testing"

	"github.com/reiver/go-ini"
)

func TestUnmarshal_errNilBytes(t *testing.T) {

	var dst map[string]string = map[string]string{}

	err := ini.Unmarshal(nil, &dst)

	if nil == err {
		t.Errorf("Expected an error but did not actually got one.")
		return
	}

	expected := "ini: nil bytes"
	actual   := err.Error()

	if expected != actual {
		t.Errorf("The actual error-message is not what was expected.")
		return
	}
}
