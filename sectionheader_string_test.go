package ini_test

import (
	"testing"

	"github.com/reiver/go-ini"
)

func TestSectionHeaderToString(t *testing.T) {

	tests := []struct{
		Name []string
		Expected string
	}{
		{

		},
		{
			Name: []string(nil),
			Expected: "",
		},
		{
			Name: []string{},
			Expected: "",
		},



		{
			Name: []string{"apple"},
			Expected:     `[apple]`+"\n",
		},
		{
			Name: []string{"apple","Banana"},
			Expected:     `[apple.Banana]`+"\n",
		},
		{
			Name: []string{"apple","Banana","CHERRY"},
			Expected:     `[apple.Banana.CHERRY]`+"\n",
		},
		{
			Name: []string{"apple","Banana","CHERRY","dAtE"},
			Expected:     `[apple.Banana.CHERRY.dAtE]`+"\n",
		},



		{
			Name: []string{"a[b","c]d","e.f"},
			Expected:     `[a\[b.c\]d.e\.f]`+"\n",
		},
	}

	for testNumber, test := range tests {

		actual := ini.SectionHeaderToString(test.Name...)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual section-header string is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d) %q", len(expected), expected)
			t.Logf("ACTUAL:   (%d) %q", len(actual), actual)
			t.Logf("NAME: %#v", test.Name)
			continue
		}
	}
}
