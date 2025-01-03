package ini

import (
	"testing"

	"reflect"
)

func TestInternalMapStringString_ForEachKeyMapValue(t *testing.T) {

	tests := []struct{
		Value internalMapStringString
		Expected []string
	}{
		{
			Expected: []string(nil),
		},
		{
			Value: internalMapStringString{},
			Expected: []string(nil),
		},
		{
			Value: internalMapStringString{nil},
			Expected: []string(nil),
		},
		{
			Value: internalMapStringString{map[string]string{}},
			Expected: []string(nil),
		},



		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE"}},
			Expected: []string(nil),
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE"}},
			Expected: []string(nil),
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
			Expected: []string(nil),
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
			Expected: []string(nil),
		},
	}

	for testNumber, test := range tests {

		var actual []string
		err := test.Value.ForEachKeyMapValue(func(key string, mapWrapper internalMapWrapper)error{
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
