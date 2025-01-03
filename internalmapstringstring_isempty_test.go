package ini

import (
	"testing"
)

func TestInternalMapStringString_IsEmpty(t *testing.T) {

	tests := []struct{
		Value internalMapStringString
		Expected bool
	}{
		{
			Expected: true,
		},
		{
			Value: internalMapStringString{},
			Expected: true,
		},
		{
			Value: internalMapStringString{nil},
			Expected: true,
		},
		{
			Value: internalMapStringString{map[string]string{}},
			Expected: true,
		},



		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE"}},
			Expected: false,
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE"}},
			Expected: false,
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
			Expected: false,
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.IsEmpty()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual is-empty is not what was expected." , testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}
