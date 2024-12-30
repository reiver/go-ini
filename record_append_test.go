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
				ToRecord(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE").
				ToRecord(),
			Name: "apple",
			Value: "TWO",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO").
				ToRecord(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO").
				ToRecord(),
			Name: "apple",
			Value: "THREE",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE").
				ToRecord(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE").
				ToRecord(),
			Name: "apple",
			Value: "FOUR",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ToRecord(),
		},



		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ToRecord(),
			Name: "Banana",
			Value: "Once",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once").
				ToRecord(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once").
				ToRecord(),
			Name: "Banana",
			Value: "Twice",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once","Twice").
				ToRecord(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice").
				ToRecord(),
			Name: "Banana",
			Value: "Thrice",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice").
				ToRecord(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice").
				ToRecord(),
			Name: "Banana",
			Value: "Fource",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ToRecord(),
		},



		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ToRecord(),
			Name: "CHERRY",
			Value: "1",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1").
				ToRecord(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1").
				ToRecord(),
			Name: "CHERRY",
			Value: "2",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2").
				ToRecord(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2").
				ToRecord(),
			Name: "CHERRY",
			Value: "3",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3").
				ToRecord(),
		},
		{
			Record:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3").
				ToRecord(),
			Name: "CHERRY",
			Value: "4",
			Expected:
				ini.NewEmptyRecord().
				ChainSet("apple", "ONE","TWO","THREE","FOUR").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3","4").
				ToRecord(),
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
