package ini

import (
	"testing"

	"reflect"
)

func TestInternalMapStringAny_ForEachKeyStringValue(t *testing.T) {

	type KeyValue struct {
		Key string
		Value string
	}

	tests := []struct{
		Value internalMapStringAny
		Expected []KeyValue
	}{
		{
			Expected: []KeyValue(nil),
		},
		{
			Value: internalMapStringAny{},
			Expected: []KeyValue(nil),
		},
		{
			Value: internalMapStringAny{nil},
			Expected: []KeyValue(nil),
		},
		{
			Value: internalMapStringAny{map[string]any{}},
			Expected: []KeyValue(nil),
		},



		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE"}},
			Expected: []KeyValue{
				KeyValue{
					Key: "apple",
					Value:"ONCE",
				},
			},
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE"}},
			Expected: []KeyValue{
				KeyValue{
					Key: "apple",
					Value:"ONCE",
				},
				KeyValue{
					Key: "Banana",
					Value:"TWICE",
				},
			},
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
			Expected: []KeyValue{
				KeyValue{
					Key: "apple",
					Value:"ONCE",
				},
				KeyValue{
					Key: "Banana",
					Value:"TWICE",
				},
				KeyValue{
					Key: "CHERRY",
					Value:"THRICE",
				},
			},
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
			Expected: []KeyValue{
				KeyValue{
					Key: "apple",
					Value:"ONCE",
				},
				KeyValue{
					Key: "Banana",
					Value:"TWICE",
				},
				KeyValue{
					Key: "CHERRY",
					Value:"THRICE",
				},
				KeyValue{
					Key: "dAtE",
					Value:"FOURCE",
				},
			},
		},



		{
			Value: internalMapStringAny{map[string]any{
				"apple":"ONCE",
				"NOPE-1":1,
				"Banana":"TWICE",
				"NOPE-2":true,
				"CHERRY":"THRICE",
				"NOPE-3":map[string]any{
					"ONE":"1",
					"TWO":"2",
					"THREE":"3",
				},
				"dAtE":"FOURCE",
			}},
			Expected: []KeyValue{
				KeyValue{
					Key: "apple",
					Value:"ONCE",
				},
				KeyValue{
					Key: "Banana",
					Value:"TWICE",
				},
				KeyValue{
					Key: "CHERRY",
					Value:"THRICE",
				},
				KeyValue{
					Key: "dAtE",
					Value:"FOURCE",
				},
			},
		},
	}

	for testNumber, test := range tests {

		var actual []KeyValue
		err := test.Value.ForEachKeyStringValue(func(key string, value string)error{
			keyvalue := KeyValue{
				Key: key,
				Value: value,
			}

			actual = append(actual, keyvalue)

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
