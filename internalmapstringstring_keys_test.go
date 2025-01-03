package ini

import (
	"testing"

	"reflect"
)

func TestInternalMapStringString_Keys(t *testing.T) {

	tests := []struct{
		Value internalMapStringString
		Expected []string
	}{
		{
			Expected: []string{},
		},
		{
			Value: internalMapStringString{},
			Expected: []string{},
		},
		{
			Value: internalMapStringString{nil},
			Expected: []string{},
		},
		{
			Value: internalMapStringString{map[string]string{}},
			Expected: []string{},
		},



		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE"}},
			Expected: []string{"apple"},
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE"}},
			Expected: []string{"apple","Banana"},
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
			Expected: []string{"apple","Banana","CHERRY"},
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
			Expected: []string{"apple","Banana","CHERRY","dAtE"},
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
