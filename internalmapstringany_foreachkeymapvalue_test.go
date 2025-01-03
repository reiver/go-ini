package ini

import (
	"testing"

	"reflect"
)

func TestInternalMapStringAny_ForEachKeyMapValue(t *testing.T) {

	tests := []struct{
		Value internalMapStringAny
		Expected []string
	}{
		{
			Expected: []string(nil),
		},
		{
			Value: internalMapStringAny{},
			Expected: []string(nil),
		},
		{
			Value: internalMapStringAny{nil},
			Expected: []string(nil),
		},
		{
			Value: internalMapStringAny{map[string]any{}},
			Expected: []string(nil),
		},



		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE"}},
			Expected: []string(nil),
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE"}},
			Expected: []string(nil),
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
			Expected: []string(nil),
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
			Expected: []string(nil),
		},



		{
			Value: internalMapStringAny{map[string]any{
				"apple":"ONCE",
				"NOPE-1":1,
				"Banana":"TWICE",
				"NOPE-2":true,
				"CHERRY":"THRICE",
				"YUP-1":map[string]any{
					"ONE":"1",
					"TWO":"2",
					"THREE":"3",
				},
				"dAtE":"FOURCE",
				"YUP-2":map[string]string{
					"YEK":"1",
					"DO":"2",
					"SE":"3",
				},
			}},
			Expected: []string{
				"YUP-1",
				"YUP-2",
			},
		},
	}

	for testNumber, test := range tests {

		var actual []string
		err := test.Value.ForEachKeyMapValue(func(key string, mapWrapper Marshaler)error{
			actual = append(actual, key)

			return nil
		})

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual for-each-key-string-value is not what was expected." , testNumber)
			t.Logf("EXPECTED: (%d) %#v", len(expected), expected)
			t.Logf("ACTUAL:   (%d) %#v", len(actual), actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}
