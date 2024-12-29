package ini

import (
	"testing"

	"reflect"

	"github.com/reiver/go-val"
)

func TestRecord_All(t *testing.T) {

	tests := []struct{
		Record Record
		Name string
		Expected []string
	}{
		{
			Name:"",
			Expected: []string{},
		},
		{
			Name:"apple",
			Expected: []string{},
		},
		{
			Name:"Banana",
			Expected: []string{},
		},
		{
			Name:"CHERRY",
			Expected: []string{},
		},
		{
			Name:"dAtE",
			Expected: []string{},
		},



		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
			Name:"",
			Expected: []string{},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
			Name:"apple",
			Expected: []string{                   "one","two","three","four"},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
			Name:"Banana",
			Expected: []string{},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
			Name:"CHERRY",
			Expected: []string{},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
			Name:"dAtE",
			Expected: []string{},
		},



		{
			Record: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
			Name:"",
			Expected: []string{},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
			Name:"apple",
			Expected: []string{},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
			Name:"Banana",
			Expected: []string{                    "Once", "Twice", "Thrice", "Fource"},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
			Name:"CHERRY",
			Expected: []string{},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
			Name:"dAtE",
			Expected: []string{},
		},



		{
			Record: Record{data:map[string]*val.Values[string]{
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"",
			Expected: []string{},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"apple",
			Expected: []string{},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"Banana",
			Expected: []string{},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"CHERRY",
			Expected: []string{                    "1","2","3","4"},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"dAtE",
			Expected: []string{},
		},



		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"",
			Expected: []string{},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"apple",
			Expected: []string{                   "one","two","three","four"},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"Banana",
			Expected: []string{                    "Once", "Twice", "Thrice", "Fource"},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"CHERRY",
			Expected: []string{                    "1","2","3","4"},
		},
		{
			Record: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
			Name:"dAtE",
			Expected: []string{},
		},
	}

	for testNumber, test := range tests {

		actual := test.Record.All(test.Name)

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual .All() results are not what were expected.", testNumber)
			t.Logf("EXPECTED: (%d) %#v", len(expected), expected)
			t.Logf("ACTUAL:   (%d) %#v", len(actual), actual)
			t.Logf("RECORD: %#v", test.Record)
			t.Logf("NAME: %q", test.Name)
			continue
		}
	}
}
