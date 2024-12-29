package ini_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-ini"
)

func TestRecord_Names(t *testing.T) {

	tests := []struct{
		Record ini.Record
		Expected []string
	}{
		{
			Expected: []string{},
		},
		{
			Record: ini.EmptyRecord(),
			Expected: []string{},
		},



		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				Record(),
			Expected: []string{"apple"},
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				Record(),
			Expected: []string{"Banana","apple"},
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3","4").
				Record(),
			Expected: []string{"Banana","CHERRY","apple"},
		},
	}

	for testNumber, test := range tests {

		actual := test.Record.Names()

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual .Names() results are not what were expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("RECORD: %#v", test.Record)
			continue
		}
	}
}
