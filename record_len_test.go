package ini

import (
	"testing"

	"reflect"

	"github.com/reiver/go-val"
)

func TestRecord_Len(t *testing.T) {

	tests := []struct{
		Record Record
		Expected int
	}{
		{
			Expected: 0,
		},



		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
			Expected: 1,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
			Expected: 1,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Expected: 1,
		},



		{
			Record: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Expected: 2,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Expected: 2,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
			Expected: 2,
		},



		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Expected: 3,
		},
	}

	for testNumber, test := range tests {

		actual := test.Record.Len()

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual .Len() results are not what were expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("RECORD: %#v", test.Record)
			continue
		}
	}
}
