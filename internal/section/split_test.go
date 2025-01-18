package inisection

import (
	"testing"

	"reflect"
)

func TestSplit(t *testing.T) {

	tests := []struct{
		Value string
		Expected []string
	}{
		{
			Value: "",
			Expected: []string{},
		},


		{
			Value: "A",
			Expected: []string{"A"},
		},
		{
			Value: "Z",
			Expected: []string{"Z"},
		},
		{
			Value: "a",
			Expected: []string{"a"},
		},
		{
			Value: "z",
			Expected: []string{"z"},
		},
		{
			Value: "۵",
			Expected: []string{"۵"},
		},



		{
			Value: "۰.۱.۲.۳.۴.۵.۶.۷.۸.۹",
			Expected: []string{"۰", "۱", "۲", "۳", "۴", "۵", "۶", "۷", "۸", "۹"},
		},
		{
			Value: "۰0.۱1.۲2.۳3.۴4.۵5.۶6.۷7.۸8.۹9",
			Expected: []string{"۰0", "۱1", "۲2", "۳3", "۴4", "۵5", "۶6", "۷7", "۸8", "۹9"},
		},




		{
			Value: "apple",
			Expected: []string{"apple"},
		},
		{
			Value: "apple.banana",
			Expected: []string{"apple", "banana"},
		},
		{
			Value: "apple.banana.cherry",
			Expected: []string{"apple", "banana", "cherry"},
		},



		{
			Value: `abcd\.efg.hi.j`,
			Expected: []string{"abcd.efg", "hi", "j"},
		},



		{
			Value: `apple\`,
			Expected: []string{"apple"},
		},
		{
			Value: `apple.banana\`,
			Expected: []string{"apple", "banana"},
		},
		{
			Value: `apple.banana.cherry\`,
			Expected: []string{"apple", "banana", "cherry"},
		},
	}

	for testNumber, test := range tests {

		actual := split(test.Value)

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual (split) section-name is not what was expected.", testNumber)
			t.Log("EXPECTED")
			for index, part := range expected {
				t.Logf("\t- [%d] %q", index, part)
			}
			t.Log("ACTUAL:")
			for index, part := range actual {
				t.Logf("\t- [%d] %q", index, part)
			}
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}
