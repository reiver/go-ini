package ini

import (
	"testing"

	"reflect"
)

func TestInternalMapStringString_ForEachKeyStringValue(t *testing.T) {

	type KeyValue struct {
		Key string
		Value string
	}

	tests := []struct{
		Value internalMapStringString
		Expected []KeyValue
	}{
		{
			Expected: []KeyValue(nil),
		},
		{
			Value: internalMapStringString{},
			Expected: []KeyValue(nil),
		},
		{
			Value: internalMapStringString{nil},
			Expected: []KeyValue(nil),
		},
		{
			Value: internalMapStringString{map[string]string{}},
			Expected: []KeyValue(nil),
		},



		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE"}},
			Expected: []KeyValue{
				KeyValue{
					Key: "apple",
					Value:"ONCE",
				},
			},
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE"}},
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
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
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
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
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
