package ini

import (
	"testing"

	"reflect"
)

func TestInternalSliceMapStringString_IsEmpty(t *testing.T) {

	tests := []struct{
		Value internalSliceMapStringString
		Expected bool
	}{
		{

			Expected: true,
		},
		{
			Value: internalSliceMapStringString{},
			Expected: true,
		},
		{
			Value: internalSliceMapStringString{nil},
			Expected: true,
		},
		{
			Value: internalSliceMapStringString{[]map[string]string(nil)},
			Expected: true,
		},
		{
			Value: internalSliceMapStringString{[]map[string]string{}},
			Expected: true,
		},



		{

			Expected: true,
		},
		{
			Value: internalSliceMapStringString{},
			Expected: true,
		},
		{
			Value: internalSliceMapStringString{nil},
			Expected: true,
		},
		{
			Value: internalSliceMapStringString{[]map[string]string(nil)},
			Expected: true,
		},
		{
			Value: internalSliceMapStringString{[]map[string]string{}},
			Expected: true,
		},



		{

			Expected: true,
		},
		{
			Value: internalSliceMapStringString{},
			Expected: true,
		},
		{
			Value: internalSliceMapStringString{nil},
			Expected: true,
		},
		{
			Value: internalSliceMapStringString{[]map[string]string(nil)},
			Expected: true,
		},
		{
			Value: internalSliceMapStringString{[]map[string]string{}},
			Expected: true,
		},



		{
			Value: internalSliceMapStringString{[]map[string]string{
				map[string]string{
					"apple":"ONCE",
					"Banana":"TWICE",
					"CHERRY":"THRICE",
					"dAtE":"FOURCE",
				},
				map[string]string{
					"apple":"1",
					"Banana":"2",
					"CHERRY":"3",
					"dAtE":"4",
				},
				map[string]string{
					"apple":"one",
					"Banana":"two",
					"CHERRY":"three",
					"dAtE":"four",
				},
			}},
			Expected: false,
		},
		{
			Value: internalSliceMapStringString{[]map[string]string{
				map[string]string{
					"apple":"ONCE",
					"Banana":"TWICE",
					"CHERRY":"THRICE",
					"dAtE":"FOURCE",
				},
				map[string]string{
					"apple":"1",
					"Banana":"2",
					"CHERRY":"3",
					"dAtE":"4",
				},
				map[string]string{
					"apple":"one",
					"Banana":"two",
					"CHERRY":"three",
					"dAtE":"four",
				},
			}},
			Expected: false,
		},
		{
			Value: internalSliceMapStringString{[]map[string]string{
				map[string]string{
					"apple":"ONCE",
					"Banana":"TWICE",
					"CHERRY":"THRICE",
					"dAtE":"FOURCE",
				},
				map[string]string{
					"apple":"1",
					"Banana":"2",
					"CHERRY":"3",
					"dAtE":"4",
				},
				map[string]string{
					"apple":"one",
					"Banana":"two",
					"CHERRY":"three",
					"dAtE":"four",
				},
			}},
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := test.Value.IsEmpty()

		{
			expected := test.Expected

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual is-empty result is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}
	}
}
