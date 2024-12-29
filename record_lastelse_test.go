package ini_test

import (
	"testing"

	"github.com/reiver/go-ini"
)

func TestRecord_LastElse(t *testing.T) {

	tests := []struct{
		Record ini.Record
		Name string
		Alternative string
		Expected string
	}{
		{
			Name: "",
			Alternative: "",
			Expected:    "",
		},
		{
			Name: "",
			Alternative: "LEFT",
			Expected:    "LEFT",
		},
		{
			Name: "",
			Alternative: "RIGHT",
			Expected:    "RIGHT",
		},



		{
			Name: "something",
			Alternative: "",
			Expected:    "",
		},
		{
			Name: "something",
			Alternative: "LEFT",
			Expected:    "LEFT",
		},
		{
			Name: "something",
			Alternative: "RIGHT",
			Expected:    "RIGHT",
		},



		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				Record(),
			Name: "",
			Alternative: "ALT",
			Expected:    "ALT",
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				Record(),
			Name: "apple",
			Alternative: "ALT",
			Expected:    "FOUR",
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				Record(),
			Name: "Banana",
			Alternative: "ALT",
			Expected:    "ALT",
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				Record(),
			Name: "CHERRY",
			Alternative: "ALT",
			Expected:    "ALT",
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				Record(),
			Name: "dAtE",
			Alternative: "ALT",
			Expected:    "ALT",
		},



		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				Record(),
			Name: "",
			Alternative: "ALT",
			Expected:    "ALT",
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				Record(),
			Name: "apple",
			Alternative: "ALT",
			Expected:    "FOUR",
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				Record(),
			Name: "Banana",
			Alternative: "ALT",
			Expected:    "Fource",
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				Record(),
			Name: "CHERRY",
			Alternative: "ALT",
			Expected:    "ALT",
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				Record(),
			Name: "dAtE",
			Alternative: "ALT",
			Expected:    "ALT",
		},



		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3","4").
				Record(),
			Name: "",
			Alternative: "ALT",
			Expected:    "ALT",
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3","4").
				Record(),
			Name: "apple",
			Alternative: "ALT",
			Expected:    "FOUR",
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3","4").
				Record(),
			Name: "Banana",
			Alternative: "ALT",
			Expected:    "Fource",
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3","4").
				Record(),
			Name: "CHERRY",
			Alternative: "ALT",
			Expected:    "4",
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3","4").
				Record(),
			Name: "dAtE",
			Alternative: "ALT",
			Expected:    "ALT",
		},
	}

	for testNumber, test := range tests {

		actual  := test.Record.LastElse(test.Name, test.Alternative)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'value' is not what was expected", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("NAME: %q", test.Name)
			t.Logf("RECORD: %#v", test.Record)
			continue
		}
	}
}
