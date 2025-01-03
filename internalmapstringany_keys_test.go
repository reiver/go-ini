package ini

import (
	"testing"

	"reflect"
)

func TestInternalMapStringAny_Keys(t *testing.T) {

	tests := []struct{
		Value internalMapStringAny
		Expected []string
	}{
		{
			Expected: []string{},
		},
		{
			Value: internalMapStringAny{},
			Expected: []string{},
		},
		{
			Value: internalMapStringAny{nil},
			Expected: []string{},
		},
		{
			Value: internalMapStringAny{map[string]any{}},
			Expected: []string{},
		},



		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE"}},
			Expected: []string{"apple"},
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE"}},
			Expected: []string{"apple","Banana"},
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
			Expected: []string{"apple","Banana","CHERRY"},
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
			Expected: []string{"apple","Banana","CHERRY","dAtE"},
		},



		{
			Value: internalMapStringAny{map[string]any{
				"apple":"ONCE",
				"Banana":"TWICE",
				"CHERRY":"THRICE",
				"dAtE":"FOURCE",
				"YUP-1":map[string]string{
					"a":"1",
					"b":"2",
					"c":"3",
					"d":"4",
				},
				"YUP-2":map[string]any{
					"A":"one",
					"B":"two",
					"C":"three",
					"D":map[string]string{
						"d-1":"aa",
						"d-2":"bb",
						"d-3":"cc",
						"d-4":"dd",
						"d-5":"ee",
					},
				},
			}},
			Expected: []string{"apple","Banana","CHERRY","dAtE","YUP-1","YUP-2"},
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.Keys()

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual keys is not what was expected." , testNumber)
			t.Logf("EXPECTED: (%d) %#v", len(expected), expected)
			t.Logf("ACTUAL:   (%d) %#v", len(actual), actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}
