package ini_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-ini"
)

func TestRecord_Append(t *testing.T) {

	tests := []struct{
		Record ini.Record
		Name string
		Value string
		Expected ini.Record
	}{
		{
			Record:
				ini.EmptyRecord(),
			Name: "apple",
			Value: "ONE",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE").
				Record(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE").
				Record(),
			Name: "apple",
			Value: "TWO",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO").
				Record(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO").
				Record(),
			Name: "apple",
			Value: "THREE",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE").
				Record(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE").
				Record(),
			Name: "apple",
			Value: "FOUR",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				Record(),
		},



		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				Record(),
			Name: "Banana",
			Value: "Once",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once").
				Record(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once").
				Record(),
			Name: "Banana",
			Value: "Twice",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once","Twice").
				Record(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice").
				Record(),
			Name: "Banana",
			Value: "Thrice",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice").
				Record(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice").
				Record(),
			Name: "Banana",
			Value: "Fource",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				Record(),
		},



		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				Record(),
			Name: "CHERRY",
			Value: "1",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1").
				Record(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1").
				Record(),
			Name: "CHERRY",
			Value: "2",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2").
				Record(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2").
				Record(),
			Name: "CHERRY",
			Value: "3",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3").
				Record(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3").
				Record(),
			Name: "CHERRY",
			Value: "4",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3","4").
				Record(),
		},
	}

	for testNumber, test := range tests {

		actual := test.Record
		actual.Append(test.Name, test.Value)

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual result of .Append() is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%#v", expected)
			t.Logf("ACTUAL:\n%#v", actual)
			t.Logf(": %q", test.Name)
			t.Logf(": %q", test.Value)
			continue
		}

	}
}
