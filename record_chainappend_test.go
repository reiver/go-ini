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
				ChainAppend("one",   "apple",).
				ChainAppend("two",   "apple",).
				ChainAppend("three", "apple").
				ChainAppend("four",  "apple").
				ToRecord(),
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
		},



		{
			Record:
				NewEmptyRecord().
				ChainAppend("one",   "apple",).
				ChainAppend("two",   "apple",).
				ChainAppend("three", "apple",).
				ChainAppend("four",  "apple",).
				ChainAppend("Once",   "Banana").
				ChainAppend("Twice",  "Banana").
				ChainAppend("Thrice", "Banana").
				ChainAppend("Fource", "Banana").
				ToRecord(),
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
		},
		{
			Record:
				NewEmptyRecord().
				ChainAppend("one", "apple").
				ChainAppend("Once", "Banana").
				ChainAppend("two", "apple").
				ChainAppend("Twice", "Banana").
				ChainAppend("three", "apple").
				ChainAppend("Thrice", "Banana").
				ChainAppend("four", "apple").
				ChainAppend("Fource", "Banana").
				ToRecord(),
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
		},



		{
			Record:
				NewEmptyRecord().
				ChainAppend("one",   "apple").
				ChainAppend("two",   "apple").
				ChainAppend("three", "apple").
				ChainAppend("four",  "apple").
				ChainAppend("Once",   "Banana").
				ChainAppend("Twice",  "Banana").
				ChainAppend("Thrice", "Banana").
				ChainAppend("Fource", "Banana").
				ChainAppend("1", "CHERRY").
				ChainAppend("2", "CHERRY").
				ChainAppend("3", "CHERRY").
				ChainAppend("4", "CHERRY").
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
				ChainAppend("1", "CHERRY").
				ChainAppend("Once", "Banana").
				ChainAppend("one", "apple").
				ChainAppend("Twice", "Banana").
				ChainAppend("two", "apple").
				ChainAppend("2", "CHERRY").
				ChainAppend("three", "apple").
				ChainAppend("3", "CHERRY").
				ChainAppend("Thrice", "Banana").
				ChainAppend("four", "apple").
				ChainAppend("Fource", "Banana").
				ChainAppend("4", "CHERRY").
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
