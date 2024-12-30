package ini

import (
	"testing"

	"reflect"

	"github.com/reiver/go-val"
)

func TestRecord_ChainAppend(t *testing.T) {

	tests := []struct{
		Record Record
		Expected Record
	}{
		{
			Record: EmptyRecord(),
			Expected: Record{},
		},



		{
			Record: NewEmptyRecord().
				ChainAppend("apple", "one").
				ChainAppend("apple", "two").
				ChainAppend("apple", "three").
				ChainAppend("apple", "four").
				ToRecord(),
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
		},



		{
			Record:
				NewEmptyRecord().
				ChainAppend("apple", "one").
				ChainAppend("apple", "two").
				ChainAppend("apple", "three").
				ChainAppend("apple", "four").
				ChainAppend("Banana", "Once").
				ChainAppend("Banana", "Twice").
				ChainAppend("Banana", "Thrice").
				ChainAppend("Banana", "Fource").
				ToRecord(),
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
		},
		{
			Record:
				NewEmptyRecord().
				ChainAppend("apple", "one").
				ChainAppend("Banana", "Once").
				ChainAppend("apple", "two").
				ChainAppend("Banana", "Twice").
				ChainAppend("apple", "three").
				ChainAppend("Banana", "Thrice").
				ChainAppend("apple", "four").
				ChainAppend("Banana", "Fource").
				ToRecord(),
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
		},



		{
			Record:
				NewEmptyRecord().
				ChainAppend("apple", "one").
				ChainAppend("apple", "two").
				ChainAppend("apple", "three").
				ChainAppend("apple", "four").
				ChainAppend("Banana", "Once").
				ChainAppend("Banana", "Twice").
				ChainAppend("Banana", "Thrice").
				ChainAppend("Banana", "Fource").
				ChainAppend("CHERRY", "1").
				ChainAppend("CHERRY", "2").
				ChainAppend("CHERRY", "3").
				ChainAppend("CHERRY", "4").
				ToRecord(),
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
		},
		{
			Record:
				NewEmptyRecord().
				ChainAppend("CHERRY", "1").
				ChainAppend("Banana", "Once").
				ChainAppend("apple", "one").
				ChainAppend("Banana", "Twice").
				ChainAppend("apple", "two").
				ChainAppend("CHERRY", "2").
				ChainAppend("apple", "three").
				ChainAppend("CHERRY", "3").
				ChainAppend("Banana", "Thrice").
				ChainAppend("apple", "four").
				ChainAppend("Banana", "Fource").
				ChainAppend("CHERRY", "4").
				ToRecord(),
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},

		},
	}

	for testNumber, test := range tests {

		actual := test.Record

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual .ChainAppend() results are not what were expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
