package ini

import (
	"testing"

	"reflect"

	"github.com/reiver/go-val"
)

func TestRecord_Unet(t *testing.T) {

	tests := []struct{
		Record Record
		Name string
		Expected Record
	}{
		{
			Record: EmptyRecord(),
			Name: "",
			Expected: Record{},
		},
		{
			Record: EmptyRecord(),
			Name: "apple",
			Expected: Record{},
		},
		{
			Record: EmptyRecord(),
			Name: "Banana",
			Expected: Record{},
		},
		{
			Record: EmptyRecord(),
			Name: "CHERRY",
			Expected: Record{},
		},
		{
			Record: EmptyRecord(),
			Name: "dAtE",
			Expected: Record{},
		},



		{
			Record: NewEmptyRecord().ChainSet("apple", "one","two","three","four").ToRecord(),
			Name: "",
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
		},
		{
			Record: NewEmptyRecord().ChainSet("apple", "one","two","three","four").ToRecord(),
			Name: "apple",
			Expected: Record{data:map[string]*val.Values[string]{
			}},
		},
		{
			Record: NewEmptyRecord().ChainSet("apple", "one","two","three","four").ToRecord(),
			Name: "Banana",
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
		},
		{
			Record: NewEmptyRecord().ChainSet("apple", "one","two","three","four").ToRecord(),
			Name: "CHERRY",
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
		},
		{
			Record: NewEmptyRecord().ChainSet("apple", "one","two","three","four").ToRecord(),
			Name: "dAtE",
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
		},



		{
			Record:
				NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ToRecord(),
			Name: "",
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
		},
		{
			Record:
				NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ToRecord(),
			Name: "apple",
			Expected: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
		},
		{
			Record:
				NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ToRecord(),
			Name: "Banana",
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
			}},
		},
		{
			Record:
				NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ToRecord(),
			Name: "CHERRY",
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
		},
		{
			Record:
				NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ToRecord(),
			Name: "dAtE",
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
		},



		{
			Record:
				NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3","4").
				ToRecord(),
			Name: "",
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
		},
		{
			Record:
				NewEmptyRecord().
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3","4").
				ToRecord(),
			Name: "apple",
			Expected: Record{data:map[string]*val.Values[string]{
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
		},
		{
			Record:
				NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("CHERRY", "1","2","3","4").
				ToRecord(),
			Name: "Banana",
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
		},
		{
			Record:
				NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ToRecord(),
			Name: "CHERRY",
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
			}},
		},
		{
			Record:
				NewEmptyRecord().
				ChainSet("apple",  "one","two","three","four").
				ChainSet("Banana", "Once", "Twice", "Thrice", "Fource").
				ChainSet("CHERRY", "1","2","3","4").
				ToRecord(),
			Name: "dAtE",
			Expected: Record{data:map[string]*val.Values[string]{
				"apple":val.NewValues[string]("one","two","three","four"),
				"Banana":val.NewValues[string]("Once", "Twice", "Thrice", "Fource"),
				"CHERRY":val.NewValues[string]("1","2","3","4"),
			}},
		},
	}

	for testNumber, test := range tests {

		actual := test.Record
		actual.Unset(test.Name)

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual .Unset() results are not what were expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("EXPECTED-KEYS: %#v", expected.Keys())
			t.Logf("ACTUAL-KEYS:   %#v", actual.Keys())
			t.Logf("NAME: %q", test.Name)
			continue
		}
	}
}
