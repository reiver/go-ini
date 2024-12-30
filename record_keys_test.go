package ini_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-ini"
)

func TestRecord_Keys(t *testing.T) {

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
				ToRecord(),
			Expected: []string{"apple"},
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ToRecord(),
			Expected: []string{"Banana","apple"},
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3","4").
				ToRecord(),
			Expected: []string{"Banana","CHERRY","apple"},
		},
	}

	for testNumber, test := range tests {

		actual := test.Record.Keys()

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual .Keys() results are not what were expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("RECORD: %#v", test.Record)
			continue
		}
	}
}
