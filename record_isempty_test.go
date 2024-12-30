package ini_test

import (
	"testing"

	"github.com/reiver/go-ini"
)

func TestRecord_IsEmpty(t *testing.T) {

	tests := []struct{
		Record ini.Record
		Expected bool
	}{
		{
			Expected: true,
		},
		{
			Record: ini.EmptyRecord(),
			Expected: true,
		},



		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("something").
				ToRecord(),
			Expected: true,
		},



		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "one","two","three","four").
				ToRecord(),
			Expected: false,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ToRecord(),
			Expected: false,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3","4").
				ToRecord(),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := test.Record.IsEmpty()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual .IsEmpty() results are not what were expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("RECORD: %#v", test.Record)
			continue
		}
	}
}
