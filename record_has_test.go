package ini

import (
	"testing"

	"reflect"

	"github.com/reiver/go-val"
)

func TestRecord_Has(t *testing.T) {

	tests := []struct{
		Record Record
		Name string
		Expected bool
	}{
		{
			Name:"",
			Expected: false,
		},
		{
			Name:"apple",
			Expected: false,
		},
		{
			Name:"Banana",
			Expected: false,
		},
		{
			Name:"CHERRY",
			Expected: false,
		},
		{
			Name:"dAtE",
			Expected: false,
		},



		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
			Name:"",
			Expected: false,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
			Name:"apple",
			Expected: true,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
			Name:"Banana",
			Expected: false,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
			Name:"CHERRY",
			Expected: false,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
			Name:"dAtE",
			Expected: false,
		},



		{
			Record: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
			Name:"",
			Expected: false,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
			Name:"apple",
			Expected: false,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
			Name:"Banana",
			Expected: true,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
			Name:"CHERRY",
			Expected: false,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
			Name:"dAtE",
			Expected: false,
		},



		{
			Record: Record{data:map[string]*val.Values[string]{
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"",
			Expected: false,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"apple",
			Expected: false,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"Banana",
			Expected: false,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"CHERRY",
			Expected: true,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"dAtE",
			Expected: false,
		},



		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"",
			Expected: false,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"apple",
			Expected: true,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"Banana",
			Expected: true,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"CHERRY",
			Expected: true,
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"dAtE",
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := test.Record.Has(test.Name)

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual .Has() results are not what were expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("RECORD: %#v", test.Record)
			t.Logf("NAME: %q", test.Name)
			continue
		}
	}
}
