package ini_test

import (
	"testing"

	"github.com/reiver/go-ini"
)

func TestUnmarshal_emptyFile(t *testing.T) {

	var bytes []byte = []byte{}

	var dst map[string]string = map[string]string{}

	err := ini.Unmarshal(bytes, &dst)
	if nil != err {
		t.Errorf("Did not expect an error, but actually gone one.")
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}

	if 0 < len(dst) {
		t.Errorf("Expect destination to be empty but actually was not.")
		return
	}
}
