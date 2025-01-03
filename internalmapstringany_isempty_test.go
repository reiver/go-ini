package ini

import (
	"testing"
)

func TestInternalMapStringAny_IsEmpty(t *testing.T) {

	tests := []struct{
		Value internalMapStringAny
		Expected bool
	}{
		{
			Expected: true,
		},
		{
			Value: internalMapStringAny{},
			Expected: true,
		},
		{
			Value: internalMapStringAny{nil},
			Expected: true,
		},
		{
			Value: internalMapStringAny{map[string]any{}},
			Expected: true,
		},



		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE"}},
			Expected: false,
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE"}},
			Expected: false,
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
			Expected: false,
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
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
