package ini_test

import (
	"testing"

	"github.com/reiver/go-ini"
)

func TestRecord_First(t *testing.T) {

	tests := []struct{
		Record ini.Record
		Name string
		ExpectedValue string
		ExpectedFound bool
	}{
		{
			ExpectedValue: "",
			ExpectedFound: false,
		},



		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				Record(),
			Name: "",
			ExpectedValue: "",
			ExpectedFound: false,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				Record(),
			Name: "apple",
			ExpectedValue: "ONE",
			ExpectedFound: true,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				Record(),
			Name: "Banana",
			ExpectedValue: "",
			ExpectedFound: false,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				Record(),
			Name: "CHERRY",
			ExpectedValue: "",
			ExpectedFound: false,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				Record(),
			Name: "dAtE",
			ExpectedValue: "",
			ExpectedFound: false,
		},



		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "once","twice","thrice","fource").
				Record(),
			Name: "",
			ExpectedValue: "",
			ExpectedFound: false,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "once","twice","thrice","fource").
				Record(),
			Name: "apple",
			ExpectedValue: "ONE",
			ExpectedFound: true,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "once","twice","thrice","fource").
				Record(),
			Name: "Banana",
			ExpectedValue: "once",
			ExpectedFound: true,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple" , "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "once","twice","thrice","fource").
				Record(),
			Name: "CHERRY",
			ExpectedValue: "",
			ExpectedFound: false,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "once","twice","thrice","fource").
				Record(),
			Name: "dAtE",
			ExpectedValue: "",
			ExpectedFound: false,
		},



		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "once","twice","thrice","fource").
				ChainSet("CHERRY", "1","2","3","4").
				Record(),
			Name: "",
			ExpectedValue: "",
			ExpectedFound: false,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "once","twice","thrice","fource").
				ChainSet("CHERRY", "1","2","3","4").
				Record(),
			Name: "apple",
			ExpectedValue: "ONE",
			ExpectedFound: true,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "once","twice","thrice","fource").
				ChainSet("CHERRY", "1","2","3","4").
				Record(),
			Name: "Banana",
			ExpectedValue: "once",
			ExpectedFound: true,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple" , "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "once","twice","thrice","fource").
				ChainSet("CHERRY", "1","2","3","4").
				Record(),
			Name: "CHERRY",
			ExpectedValue: "1",
			ExpectedFound: true,
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple",  "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "once","twice","thrice","fource").
				ChainSet("CHERRY", "1","2","3","4").
				Record(),
			Name: "dAtE",
			ExpectedValue: "",
			ExpectedFound: false,
		},
	}

	for testNumber, test := range tests {

		actualValue, actualFound := test.Record.First(test.Name)

		{
			actual := actualFound
			expected := test.ExpectedFound

			if expected != actual {
				t.Errorf("For test #%d, the actual 'found' is not what was expected", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("NAME: %q", test.Name)
				t.Logf("RECORD: %#v", test.Record)
				continue
			}
		}

		{
			actual := actualValue
			expected := test.ExpectedValue

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
}
